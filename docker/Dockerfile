FROM ubuntu:20.04

# Docker ENV commands with secrets
ENV PASSWORD=MySecurePassword123
ENV SECRET_KEY=AnotherSecretValue456
ENV API_TOKEN=AbCdEf1234567890ghijkl
ENV DB_PASSWORD=DatabasePass2024!

# Docker RUN commands with secrets
RUN TRIVY_PASSWORD=PVPkantek2025#
RUN apt-get update && apt-get install -y curl

# More environment variables
ENV APP_SECRET=MyAppSecret12345
ENV ENCRYPTION_KEY=EncryptKey9876543
ENV ADMIN_PASSWORD=AdminPassword2024
ENV USER_PASSWORD=UserPass789012
ENV AUTH_SECRET=AuthSecret345678
ENV SERVICE_KEY=ServiceKey901234
ENV BACKUP_PASSWORD=BackupPass567890
ENV CACHE_SECRET=CacheSecret123456

WORKDIR /app
COPY . /app

CMD ["python", "app.py"]

