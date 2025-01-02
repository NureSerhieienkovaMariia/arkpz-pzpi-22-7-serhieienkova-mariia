#include <OneWire.h>
#include <DallasTemperature.h>
#include <WiFi.h>
#include <PubSubClient.h>

// Налаштування датчика температури
#define ONE_WIRE_BUS 23 // Пін для DS18B20
OneWire oneWire(ONE_WIRE_BUS);
DallasTemperature sensors(&oneWire);

// Налаштування генератора пульсу
#define PULSE_PIN 2            // Пін генератора пульсу (GPIO2)
#define SAMPLING_INTERVAL 1000 // Інтервал для вибірки пульсу (в мілісекундах)
volatile uint16_t pulse = 0;
uint16_t count = 0;
int heartRate = 0;
unsigned long lastPulseTime = 0;

// Налаштування WiFi
const char* ssid = "Wokwi-GUEST"; // Ім'я мережі WiFi
const char* password = "";

// Налаштування MQTT
const char* mqtt_server = "broker.hivemq.com"; // Адреса брокера MQTT
const int mqtt_port = 1883;                   // Порт MQTT
const char* mqtt_topic = "agewell/data";      // Тема для публікації даних
const char* mqtt_user = "";
const char* mqtt_password = "";

WiFiClient espClient;
PubSubClient client(espClient);

// Порогові значення для температури та тиску
const float temperatureThresholdHigh = 41.0;  // Максимальна температура
const float temperatureThresholdLow = 28.0;   // Мінімальна температура
const float pressureThresholdHigh = 180.0;    // Максимальний тиск
const float pressureThresholdLow = 110.0;     // Мінімальний тиск

const int deviceId = 1; // Ідентифікатор пристрою

// Функція для обробки імпульсів пульсу
void HeartRateInterrupt() {
  pulse++;  // Збільшуємо лічильник пульсу при кожному імпульсі
}

void connectToWiFi() {
  Serial.print("Connecting to WiFi");
  WiFi.begin(ssid, password);
  while (WiFi.status() != WL_CONNECTED) {
    Serial.print(".");
  }
  Serial.println("Connected to WiFi");
}

void connectToMQTT() {
  while (!client.connected()) {
    Serial.print("Connecting to MQTT...");
    if (client.connect("ESP32Client", mqtt_user, mqtt_password)) {
      Serial.println("Connected to MQTT");
    } else {
      Serial.print("Failed, rc=");
      Serial.print(client.state());
      Serial.println(" trying again in 5 seconds");
      delay(5000);
    }
  }
}

void setup() {
  // Ініціалізація серійного порту
  Serial.begin(115200);

  // Ініціалізація датчика температури
  sensors.begin();

  // Ініціалізація генератора пульсу
  pinMode(PULSE_PIN, INPUT);
  attachInterrupt(digitalPinToInterrupt(PULSE_PIN), HeartRateInterrupt, RISING); // Налаштування переривання на пін PULSE_PIN

  // Підключення до WiFi
  connectToWiFi();

  // Налаштування MQTT
  client.setServer(mqtt_server, mqtt_port);
  connectToMQTT();

  Serial.println("Setup complete!\n");
}

void loop() {
  if (!client.connected()) {
    connectToMQTT();
  }
  client.loop();

  // Зчитування температури
  sensors.requestTemperatures();
  float temperature = sensors.getTempCByIndex(0);

  // Вимірювання пульсу за 1 секунду
  unsigned long currentMillis = millis();
  if (currentMillis - lastPulseTime >= SAMPLING_INTERVAL) {
    heartRate = pulse * 60;  // Перетворюємо кількість імпульсів за 1 секунду в кількість ударів за хвилину
    pulse = 0;              // Скидаємо лічильник імпульсів
    lastPulseTime = currentMillis; // Оновлюємо час

    // Генерація даних для систолічного та діастолічного тиску
    int systolicPressure = random(110, 180);  // Генерація випадкового систолічного тиску
    int diastolicPressure = random(70, 120);  // Генерація випадкового діастолічного тиску

    String healthStatus = "Normal";

    // Перевірка стану за температурою
    if (temperature >= temperatureThresholdHigh) {
      healthStatus = "Patient in critical condition due to high temperature!";
    } else if (temperature <= temperatureThresholdLow) {
      healthStatus = "Patient in critical condition due to low temperature!";
    }

    // Перевірка стану за тиском
    if (systolicPressure >= pressureThresholdHigh || diastolicPressure >= pressureThresholdLow) {
      if (healthStatus == "Normal") {
        healthStatus = "Patient in critical condition due to pressure!";
      } else {
        healthStatus += " And critical condition due to pressure!";
      }
    }

    // Створення JSON-рядка для відправлення
    String payload = "{\"temperature\": " + String(temperature) +
                      ", \"systolic_blood_pressure\":" + String(systolicPressure) +
                      ",\"distolic_blood_pressure\":" + String(diastolicPressure) +
                      ", \"pulse\": " + String(heartRate) +
                      ", \"device_id\": " + String(deviceId) + "}";

    // Публікація даних у MQTT
    if (temperature != DEVICE_DISCONNECTED_C) {
      client.publish(mqtt_topic, payload.c_str());
      Serial.println("Published to MQTT: " + payload);
      Serial.println("Patient Health Status: " + healthStatus);
    } else {
      Serial.println("Error: Could not read temperature data!");
    }
  }

  // Затримка на 5 секунд
  delay(5000);
}