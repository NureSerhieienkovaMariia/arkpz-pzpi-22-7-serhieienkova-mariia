# Код до рефакторингу
#include <iostream>
#include <vector>
#include <string>
using namespace std;

class Employee {
public:
    string type; // Тип працівника (менеджер, розробник, стажер)

    Employee(string type) : type(type) {}

    // Обчислення зарплати залежно від типу працівника
    double calculateSalary(int hoursWorked) {
        if (type == "manager") {
            return hoursWorked * 50;
        }
        else if (type == "developer") {
            return hoursWorked * 40;
        }
        else {
            return hoursWorked * 30;
        }
    }
};

class Order {
public:
    vector<string> items; // Список товарів у замовленні
    bool isProcessed = false; // Статус обробки замовлення

    Order(vector<string> items) : items(items) {}

    // Обробка замовлення
    int processOrder() {
        if (items.empty()) {
            return -1; // Помилка: замовлення порожнє
        }
        isProcessed = true;
        return items.size();
    }
};

int main() {
    Employee manager("manager");
    Employee developer("developer");
    Employee intern("intern");

    cout << "Manager Salary: " << manager.calculateSalary(160) << endl;
    cout << "Developer Salary: " << developer.calculateSalary(160) << endl;
    cout << "Intern Salary: " << intern.calculateSalary(160) << endl;

    vector<string> items = {};
    Order order(items);

    int result = order.processOrder();
    if (result == -1) {
        cout << "Error: Order is empty!" << endl;
    }
    else {
        cout << "Order processed: " << result << " items." << endl;
    }

    return 0;
}


# Код після рефакторингу
#include <iostream>
#include <vector>
#include <string>
#include <memory>
#include <stdexcept>
using namespace std;

// Стратегія для обчислення зарплати
class SalaryStrategy {
public:
    virtual double calculate(int hoursWorked) const = 0;
    virtual ~SalaryStrategy() = default;
};

class ManagerSalaryStrategy : public SalaryStrategy {
public:
    double calculate(int hoursWorked) const override {
        return hoursWorked * 50;
    }
};

class DeveloperSalaryStrategy : public SalaryStrategy {
public:
    double calculate(int hoursWorked) const override {
        return hoursWorked * 40;
    }
};

class InternSalaryStrategy : public SalaryStrategy {
public:
    double calculate(int hoursWorked) const override {
        return hoursWorked * 30;
    }
};

class Employee {
protected:
    shared_ptr<SalaryStrategy> salaryStrategy;

public:
    Employee(shared_ptr<SalaryStrategy> strategy)
        : salaryStrategy(strategy) {}

    double calculateSalary(int hoursWorked) {
        return salaryStrategy->calculate(hoursWorked);
    }
};

class EmptyOrderException : public runtime_error {
public:
    explicit EmptyOrderException(const string& message)
        : runtime_error(message) {}
};

class Order {
protected:
    vector<string> items;

public:
    bool isProcessed = false;

    Order(vector<string> items) : items(items) {}

    void processOrder() {
        if (items.empty()) {
            throw EmptyOrderException("Order cannot be empty!");
        }
        isProcessed = true;
    }

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

    try {
        vector<string> items = {};
        Order order(items);

        // Спроба обробити замовлення
        order.processOrder();
        cout << "Order processed: " << order.getItemCount() << " items." << endl;
    }
    catch (const EmptyOrderException& e) {
        cout << "Error: " << e.what() << endl;
    }

    return 0;
}
