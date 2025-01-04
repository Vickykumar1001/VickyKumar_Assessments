from flask import Flask
from flask_sqlalchemy import SQLAlchemy
from config import Config

db = SQLAlchemy()

def create_app():
    app = Flask(__name__)
    app.config.from_object(Config)
    
    db.init_app(app)

    from .routes.index_routes import index_bp as index_blueprint
    app.register_blueprint(index_blueprint)
    
    from .routes.book_routes import book as book_blueprint
    app.register_blueprint(book_blueprint)

    with app.app_context():
        db.create_all()
        return app