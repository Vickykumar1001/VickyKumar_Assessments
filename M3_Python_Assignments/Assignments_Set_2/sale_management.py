from customer_management import Customer
class Transaction(Customer):
    # inherit customer
    def __init__(self, name, email, phone, book_title, quantity_sold):
        super().__init__(name, email, phone)
        self.book_title = book_title
        self.quantity_sold = quantity_sold
    def record_transaction(self):
        return f"Customer: {self.name}, Book: {self.book_title}, Quantity Sold: {self.quantity_sold}"

def sell_book(books_list, sales_list, customer_name, book_title, quantity):
    for book in books_list:
        if book.title == book_title:
            if book.quantity >= quantity:
                book.quantity -= quantity
                sales_list.append(Transaction(customer_name, "", "", book_title, quantity))
                return f"Sale successful..! Remaining quantity is: {book.quantity}"
            else:
                return f"Error: Only {book.quantity} copies available."
    return "Error: Book not found."

def view_sales(sales_list):
    if not sales_list:
        return ["No sales records available."]
    return [sale.record_transaction() for sale in sales_list]
