// User.swift
// Модель даних користувача

import Foundation

struct User: Identifiable {
    let id: Int
    let name: String
    let email: String
    var isActive: Bool

    /// Повертає деталі користувача у вигляді форматованого рядка
    func userDetails() -> String {
        return "ID: \(id), Name: \(name), Email: \(email), Active: \(isActive)"
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
        print("User added: \(user.name)")
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
}

// Main.swift
// Основна точка входу програми

import Foundation

let manager = UserManager()
let user1 = User(id: 1, name: "Mariia", email: "mariia@example.com", isActive: true)
let user2 = User(id: 2, name: "Sophia", email: "sophia@example.com", isActive: false)

manager.addUser(user1)
manager.addUser(user2)
manager.listAllUsers()
manager.removeUser(by: 1)
manager.listAllUsers()

