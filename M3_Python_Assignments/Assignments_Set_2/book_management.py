class Book:
    def __init__(self, title, author, price, quantity):
        self.title = title
        self.author = author
        self.price = price
        self.quantity = quantity
    def display_details(self):
        return f"Title: {self.title}, Author: {self.author}, Price: {self.price}, Quantity: {self.quantity} "
    
def add_book(books_list, title, author, price, quantity):
    new_book = Book(title, author, price, quantity)
    books_list.append(new_book)
    return "Book added..!"

def view_books(books_list):
    if not books_list:
        return ["No books available."]
    return [book.display_details() for book in books_list]

def search_book(books_list, title):
    for book in books_list:
        if book.title.lower()==title.lower():
            return book.display_details()
    return ["Book not found..!"]