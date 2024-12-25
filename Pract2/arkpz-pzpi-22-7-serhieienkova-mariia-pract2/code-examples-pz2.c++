# Код до рефакторингу
#include <iostream>
#include <vector>
#include <string>
using namespace std;

class Employee {
public:
    string type;  // Тип працівника (менеджер, розробник, стажер)

    Employee(string type) : type(type) {}

    // Обчислення зарплати залежно від типу працівника
    double calculateSalary(int hoursWorked) {
        if (type == "manager") {
            return hoursWorked * 50;
        } else if (type == "developer") {
            return hoursWorked * 40;
        } else {
            return hoursWorked * 30;
        }
    }
};

class Order {
public:
    vector<string> items;  // Список товарів у замовленні
    bool isProcessed = false;  // Статус обробки замовлення

    Order(vector<string> items) : items(items) {}

    // Обробка замовлення
    int processOrder() {
        isProcessed = true;
        return items.size();  // Повертаємо кількість товарів
    }
};

int main() {
    Employee manager("manager");
    Employee developer("developer");
    Employee intern("intern");

    cout << "Manager Salary: " << manager.calculateSalary(160) << endl;
    cout << "Developer Salary: " << developer.calculateSalary(160) << endl;
    cout << "Intern Salary: " << intern.calculateSalary(160) << endl;

    vector<string> items = {"item1", "item2", "item3"};
    Order order(items);
    cout << "Order processed: " << order.processOrder() << " items." << endl;

    return 0;
}


# Код після рефакторингу
#include <iostream>
#include <vector>
#include <string>
#include <memory>
using namespace std;

// Стратегія для обчислення зарплати
class SalaryStrategy {
public:
    virtual double calculate(int hoursWorked) const = 0;
    virtual ~SalaryStrategy() = default;
};

// Стратегія для менеджера
class ManagerSalaryStrategy : public SalaryStrategy {
public:
    double calculate(int hoursWorked) const override {
        return hoursWorked * 50;
    }
};

// Стратегія для розробника
class DeveloperSalaryStrategy : public SalaryStrategy {
public:
    double calculate(int hoursWorked) const override {
        return hoursWorked * 40;
    }
};

// Стратегія для стажера
class InternSalaryStrategy : public SalaryStrategy {
public:
    double calculate(int hoursWorked) const override {
        return hoursWorked * 30;
    }
};

// Клас працівника
class Employee {
protected:
    shared_ptr<SalaryStrategy> salaryStrategy;

public:
    Employee(shared_ptr<SalaryStrategy> strategy)
        : salaryStrategy(strategy) {}

    // Обчислення зарплати за допомогою стратегії
    double calculateSalary(int hoursWorked) {
        return salaryStrategy->calculate(hoursWorked);
    }
};

// Клас замовлення
class Order {
protected:
    vector<string> items;

public:
    bool isProcessed = false;

    Order(vector<string> items) : items(items) {}

    // Обробка замовлення
    void processOrder() {
        isProcessed = true;
    }

    // Отримання кількості товарів у замовленні
    int getItemCount() const {
        return items.size();
    }
};

int main() {
    // Створення об'єктів працівників з різними стратегіями
    Employee manager(make_shared<ManagerSalaryStrategy>());
    Employee developer(make_shared<DeveloperSalaryStrategy>());
    Employee intern(make_shared<InternSalaryStrategy>());

    cout << "Manager Salary: " << manager.calculateSalary(160) << endl;
    cout << "Developer Salary: " << developer.calculateSalary(160) << endl;
    cout << "Intern Salary: " << intern.calculateSalary(160) << endl;

    vector<string> items = {"item1", "item2", "item3"};
    Order order(items);

    // Обробка замовлення
    order.processOrder();
    cout << "Order processed: " << order.getItemCount() << " items." << endl;

    return 0;
}
