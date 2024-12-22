// User.swift
// Модель даних користувача з реалізацією протоколу Identifiable

import Foundation

struct User: Identifiable {
    let id: Int
    let firstName: String
    let lastName: String
    let email: String
    var isActive: Bool

    /// Повертає повне ім'я користувача
    func fullName() -> String {
        return "\(firstName) \(lastName)"
    }
    
    /// Повертає деталі користувача у вигляді форматованого рядка
    func userDetails() -> String {
        return "ID: \(id), Name: \(fullName()), Email: \(email), Active: \(isActive)"
    }
}

// UserManager.swift
// Логіка управління списком користувачів

import Foundation

class UserManager {
    private var users: [User] = []

    /// Додає нового користувача до списку
    func addUser(_ user: User) {
        users.append(user)
        print("User added: \(user.fullName())")
    }
    
    /// Видаляє користувача за ID
    func removeUser(by id: Int) {
        users.removeAll { $0.id == id }
        print("User with ID \(id) removed")
    }

    /// Показує всіх користувачів
    func listAllUsers() {
        users.forEach { print($0.userDetails()) }
    }

    /// Функція, що перевіряє активність користувача
    func checkActiveStatus(for user: User) {
        if user.isActive {
            print("\(user.fullName()) is active.")
        } else {
            print("\(user.fullName()) is not active.")
        }
    }
}

// Discount.swift
// Функції для обчислення знижок та ціни

import Foundation

/// Перевіряє, чи є знижка дійсною
func isDiscountValid(_ discountRate: Double) -> Bool {
    return discountRate > 0 && discountRate <= 1
}

/// Обчислює суму знижки на основі ціни та ставки знижки
func calculateDiscount(price: Double, discountRate: Double) -> Double {
    return price * discountRate
}

/// Обчислює нову ціну після застосування знижки
func calculateDiscountedPrice(price: Double, discountRate: Double) -> Double {
    if isDiscountValid(discountRate) {
        let discount = calculateDiscount(price: price, discountRate: discountRate)
        return price - discount
    } else {
        print("Invalid discount rate")
        return price
    }
}

// UserTests.swift
// Тести для класу User та Discount

import XCTest

class UserTests: XCTestCase {
    
    func testUserFullName() {
        let user = User(id: 1, firstName: "Mariia", lastName: "Serhieienkova", email: "mariia@example.com", isActive: true)
        XCTAssertEqual(user.fullName(), "Mariia Serhieienkova")
    }
    
    func testUserDetails() {
        let user = User(id: 2, firstName: "Sophia", lastName: "Gorb", email: "sophia@example.com", isActive: false)
        XCTAssertEqual(user.userDetails(), "ID: 2, Name: Sophia Gorb, Email: sophia@example.com, Active: false")
    }
    
    func testDiscountValidation() {
        XCTAssertTrue(isDiscountValid(0.1))
        XCTAssertFalse(isDiscountValid(1.5))
        XCTAssertFalse(isDiscountValid(0))
    }
    
    func testDiscountedPrice() {
        let price = 100.0
        let discountRate = 0.2
        let discountedPrice = calculateDiscountedPrice(price: price, discountRate: discountRate)
        XCTAssertEqual(discountedPrice, 80.0)
        
        let invalidDiscountRate = 1.5
        let invalidDiscountedPrice = calculateDiscountedPrice(price: price, discountRate: invalidDiscountRate)
        XCTAssertEqual(invalidDiscountedPrice, 100.0) // should return original price
    }
}

// Main.swift
// Основна точка входу програми

import Foundation

let manager = UserManager()
let user1 = User(id: 1, firstName: "Mariia", lastName: "Serhieienkova", email: "mariia@example.com", isActive: true)
let user2 = User(id: 2, firstName: "Sophia", lastName: "Gorb", email: "sophia@example.com", isActive: false)

manager.addUser(user1)
manager.addUser(user2)
manager.listAllUsers()
manager.removeUser(by: 1)
manager.listAllUsers()

// Приклад використання знижок
let price = 200.0
let discountRate = 0.1
let finalPrice = calculateDiscountedPrice(price: price, discountRate: discountRate)
print("Final price after discount: \(finalPrice)")
