from app import create_app, db
from app.models import Book

def populate_db():
    books = [
        {"title": "The Great Gatsby", "author": "F. Scott Fitzgerald", "published_year": 1925, "genre": "Fiction"},
        {"title": "To Kill a Mockingbird", "author": "Harper Lee", "published_year": 1960, "genre": "Fiction"},
        {"title": "1984", "author": "George Orwell", "published_year": 1949, "genre": "Sci-Fi"},
        {"title": "Dune", "author": "Frank Herbert", "published_year": 1965, "genre": "Sci-Fi"},
        {"title": "The God of Small Things", "author": "Arundhati Roy", "published_year": 1997, "genre": "Fiction"}
    ]
    app = create_app()
    with app.app_context():        
        db.create_all()
        for book_data in books:
            book = Book(**book_data)
            db.session.add(book)
        
        db.session.commit()
        print(f"Database populated..!")

if __name__ == "__main__":
    populate_db()
