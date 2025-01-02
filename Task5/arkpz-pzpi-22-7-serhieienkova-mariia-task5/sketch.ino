#include <OneWire.h>
#include <DallasTemperature.h>
#include <WiFi.h>
#include <PubSubClient.h>

// Настройки датчика температуры
#define ONE_WIRE_BUS 23 // Пин для DS18B20
OneWire oneWire(ONE_WIRE_BUS);
DallasTemperature sensors(&oneWire);

// Pulse Generator Settings
#define PULSE_PIN 2            // Pulse generator pin (GPIO2)
#define SAMPLING_INTERVAL 1000 // Интервал для выборки пульса
volatile uint16_t pulse = 0;
uint16_t count = 0;
int heartRate = 0;
unsigned long lastPulseTime = 0; // Переменная для отслеживания времени последнего измерения пульса

// WiFi настройки
const char* ssid = "Wokwi-GUEST"; // Укажите ваш SSID
const char* password = ""; // Укажите ваш пароль

// MQTT настройки
const char* mqtt_server = "broker.hivemq.com"; // Адрес брокера
const int mqtt_port = 1883; // Порт MQTT
const char* mqtt_topic = "agewell/data"; // Тема для публикации данных
const char* mqtt_user = ""; // Пользователь MQTT (если есть)
const char* mqtt_password = ""; // Пароль MQTT (если есть)

WiFiClient espClient;
PubSubClient client(espClient);

// Пороговые значения для температуры и давления
const float temperatureThresholdHigh = 41.0;
const float temperatureThresholdLow = 28.0; 
const float pressureThresholdHigh = 180.0;
const float pressureThresholdLow = 110.0;

const int deviceId = 1;

// Функция для обработки импульсов пульса
void HeartRateInterrupt() {
  pulse++;  // Увеличиваем счетчик пульса при каждом импульсе
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
  // Инициализация последовательного порта
  Serial.begin(115200);

  // Инициализация датчика температуры
  sensors.begin();

  // Инициализация датчика пульса
  pinMode(PULSE_PIN, INPUT);
  attachInterrupt(digitalPinToInterrupt(PULSE_PIN), HeartRateInterrupt, RISING); // Настройка прерывания на пин PULSE_PIN

  // Подключение к WiFi
  connectToWiFi();

  // Настройка MQTT
  client.setServer(mqtt_server, mqtt_port);
  connectToMQTT();

  Serial.println("Setup complete!\n");
}

void loop() {
  if (!client.connected()) {
    connectToMQTT();
  }
  client.loop();

  // Считывание температуры
  sensors.requestTemperatures();
  float temperature = sensors.getTempCByIndex(0);

  // Измерение пульса за 1 секунду
  unsigned long currentMillis = millis();
  if (currentMillis - lastPulseTime >= SAMPLING_INTERVAL) {
    heartRate = pulse * 60;  // Умножаем количество импульсов за 1 секунду на 60, чтобы получить количество пульсов в минуту
    pulse = 0;  // Сбрасываем счетчик пульсов
    lastPulseTime = currentMillis;  // Обновляем время

    // Генерация данных для систолического и диастолического давления
    int systolicPressure = random(110, 180);  // Генерация случайного систолического давления
    int diastolicPressure = random(70, 120);  // Генерация случайного диастолического давления

    String healthStatus = "Normal";
    
    // Проверка состояния по температуре
    if (temperature >= temperatureThresholdHigh) {
      healthStatus = "Patient in critical condition due to high temperature!";
    } else if (temperature <= temperatureThresholdLow) {
      healthStatus = "Patient in critical condition due to low temperature!";
    }

    // Проверка состояния по давлению
    if (systolicPressure >= pressureThresholdHigh || diastolicPressure >= pressureThresholdLow) {
    if (healthStatus == "Normal") {
      healthStatus = "Patient in critical condition due to pressure!";
    } else {
      healthStatus += " And critical condition due to pressure!";
    }
}

    // Создание JSON строки для отправки
    String payload = "{\"temperature\": " + String(temperature) + 
                      ", \"systolic_blood_pressure\":" + String(systolicPressure) + 
                      ",\"distolic_blood_pressure\":" + String(diastolicPressure) +
                      ", \"pulse\": " + String(heartRate) + 
                      ", \"device_id\": " + String(deviceId) + "}";

    // Публикация данных в MQTT
    if (temperature != DEVICE_DISCONNECTED_C) {
      client.publish(mqtt_topic, payload.c_str());
      Serial.println("Published to MQTT: " + payload);
      Serial.println("Patient Health Status: " + healthStatus);
    } else {
      Serial.println("Error: Could not read temperature data!");
    }
  }

  // Задержка в 5 секунд
  delay(5000);
}
