from typing import TYPE_CHECKING

from sqlalchemy import ForeignKey
from sqlalchemy.orm import mapped_column, Mapped, relationship

from app.models.base_model import Base, UUIDMixin, TimestampsMixin

if TYPE_CHECKING:
    from app.models.user_model import User
    from app.models.comment_model import Comment


class Task(Base, UUIDMixin, TimestampsMixin):
    __tablename__ = "tasks"

    id: Mapped[int] = mapped_column(primary_key=True, index=True)
    title: Mapped[str]
    description: Mapped[str]
    user_id: Mapped[int] = mapped_column(ForeignKey("users.id"), index=True)

    # Relationship
    user: Mapped["User"] = relationship(back_populates="tasks")
    comments: Mapped[list["Comment"]] = relationship(back_populates="task")
