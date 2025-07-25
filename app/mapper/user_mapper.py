from sqlalchemy.orm import Session
from app.models.user import User

class UserMapper:
    @staticmethod
    def get_user_by_id(db: Session, user_id: int):
        return db.query(User).filter(User.id == user_id).first()

    @staticmethod
    def create_user(db: Session, name: str, email: str):
        user = User(name=name, email=email)
        db.add(user)
        db.commit()
        db.refresh(user)
        return user
