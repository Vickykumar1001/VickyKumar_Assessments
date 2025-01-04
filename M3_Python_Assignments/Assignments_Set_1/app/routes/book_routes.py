from flask import Blueprint, request, jsonify
from .. import db
from ..models import Book

book =Blueprint('book' ,__name__)

def validate_book_data(data):
    required_fields = ['title', 'author', 'published_year', 'genre']
    for field in required_fields:
        if field not in data:
            return False, f"Missing field: {field}"
    if not isinstance(data['published_year'], int) or data['published_year'] <= 0:
        return False, "Invalid published_year: must be a positive integer."
    valid_genres = ["Fiction", "Non-Fiction", "Mystery" , "Sci-Fi"]
    if data['genre'] not in valid_genres:
        return False, f"Invalid genre: must be one of {valid_genres}."
    return True, None

@book.route('/books', methods=['GET'])
def get_books():
    genre = request.args.get('genre')
    author = request.args.get('author')
    query = Book.query
    if genre:
        query = query.filter_by(genre=genre)
    if author:
        query = query.filter_by(author=author)

    books = query.all()
    return jsonify([book.as_dict() for book in books]), 200

@book.route('/books/<int:book_id>', methods=['GET'])
def get_book(book_id):
    book = Book.query.get(book_id)
    if not book:
        return jsonify({"error":"Book not found", "message":"No book exists with the provided ID"}), 404
    return jsonify(book.as_dict()), 200

@book.route('/books', methods=['POST'])
def add_book():
    data = request.get_json()
    valid, error_message = validate_book_data(data)
    if not valid:
        return jsonify({"error":"Invalid data", "message": error_message}), 400

    new_book = Book(
        title=data['title'],
        author=data['author'],
        published_year=data['published_year'],
        genre=data['genre']
    )
    db.session.add(new_book)
    db.session.commit()
    return jsonify({"message": "Book added successfully", "book_id": new_book.id}), 201


@book.route('/books/<int:book_id>', methods=['PUT'])
def update_book(book_id):
    data = request.get_json()
    valid, error_message = validate_book_data(data)
    if not valid:
        return jsonify({"error": "Invalid data", "message": error_message}), 400

    book = Book.query.get(book_id)
    if not book:
        return jsonify({"error": "Book not found", "message":"No book exists with the provided ID"}), 404

    book.title = data['title']
    book.author = data['author']
    book.published_year = data['published_year']
    book.genre = data['genre']
    db.session.commit()

    return jsonify({"message": "Book updated successfully.."}), 200

@book.route('/books/<int:book_id>', methods=['DELETE'])
def delete_book(book_id):
    book = Book.query.get(book_id)
    if not book:
        return jsonify({"error": "Book not found", "message":"No book exists with the provided ID"}), 404

    db.session.delete(book)
    db.session.commit()
    return jsonify({"message": "Book deleted successfully"}), 200
