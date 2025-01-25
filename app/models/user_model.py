import re
from typing import TYPE_CHECKING

from sqlalchemy.orm import mapped_column, Mapped, validates, relationship

from app.models.base_model import Base, UUIDMixin, TimestampMixin

if TYPE_CHECKING:
    from app.models.task_model import Task
    from app.models.comment_model import Comment

class User(Base, UUIDMixin, TimestampMixin):
    __tablename__ = "users"

    id: Mapped[int] = mapped_column(primary_key=True, index=True)
    name: Mapped[str]
    email: Mapped[str] = mapped_column(unique=True)
    password: Mapped[str]

    tasks: Mapped[list["Task"]] = relationship(back_populates="user")
    comment: Mapped[list["Comment"]] = relationship(back_populates="user")

    @validates
    def validate_email(self, key, email):
        email_regex = r'^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$'
        if not re.match(email_regex, email):
            raise ValueError("Invalid email format")
        return email.lower()
