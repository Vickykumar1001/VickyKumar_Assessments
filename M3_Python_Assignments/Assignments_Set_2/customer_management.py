class Customer:
    def __init__(self, name, email, phone):
        self.name = name
        self.email = email
        self.phone = phone
    def display_details(self):
        return f"Name: {self.name}, Email: {self.email}, Phone: {self.phone}"

def add_customer(customers_list, name, email, phone):
    new_customer = Customer(name, email, phone)
    customers_list.append(new_customer)
    return "Customer added successfully..!"

def view_customers(customers_list):
    if not customers_list:
        return ["No customers available."]
    return [customer.display_details() for customer in customers_list]
