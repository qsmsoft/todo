import datetime
import uuid as uuid_
from typing import Optional

from sqlalchemy import BIGINT, func, DateTime
from sqlalchemy.dialects.postgresql import UUID
from sqlalchemy.orm import DeclarativeBase, mapped_column, Mapped


class Base(DeclarativeBase):
    type_annotation_map = {
        int: BIGINT
    }


class UUIDMixin:
    uuid: Mapped[uuid_.UUID] = mapped_column(UUID(as_uuid=True), server_default=func.gen_random_uuid())


class TimestampMixin:
    created_at: Mapped[datetime.datetime] = mapped_column(DateTime(timezone=True), server_default=func.now())
    updated_at: Mapped[datetime.datetime] = mapped_column(DateTime(timezone=True), server_default=func.now())


class TimestampsMixin:
    created_at: Mapped[datetime.datetime] = mapped_column(DateTime(timezone=True), server_default=func.now())
    updated_at: Mapped[datetime.datetime] = mapped_column(DateTime(timezone=True), server_default=func.now())
    deleted_at: Mapped[Optional[datetime.datetime]] = mapped_column(DateTime(timezone=True))