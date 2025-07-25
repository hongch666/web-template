from sqlalchemy.orm import Session
from app.mapper.user_mapper import UserMapper

class UserService:
    def __init__(self, db: Session):
        self.db = db

    def get_user(self, user_id: int):
        return UserMapper.get_user_by_id(self.db, user_id)

    def create_user(self, name: str, email: str):
        return UserMapper.create_user(self.db, name, email)
