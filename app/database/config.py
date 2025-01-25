from contextlib import asynccontextmanager

from app.core.settings import Settings
from app.models.base_model import Base
from fastapi_cli.cli import logger
from sqlalchemy.ext.asyncio import create_async_engine, AsyncSession, async_sessionmaker, AsyncEngine

DATABASE_URL = Settings.database_url()

engine: AsyncEngine = create_async_engine(DATABASE_URL, echo=True, future=True, pool_pre_ping=True)

async_session = async_sessionmaker(bind=engine, expire_on_commit=False, class_=AsyncSession)


async def init_db():
    async with engine.begin() as conn:
        logger.info("Initializing the database...")
        await conn.run_sync(Base.metadata.create_all)
        logger.info("Database initialized successfully.")


@asynccontextmanager
async def get_async_session() -> AsyncSession:
    async with async_session() as session:
        try:
            yield session
            await session.commit()
        except Exception:
            await session.rollback()
            raise
