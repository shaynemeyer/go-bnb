# Create Database

Ensure you have Postgres installed.

---

## Create a User

```shell
psql -c "create user bookings_service;"
psql -c "alter user bookings_service with encrypted password 'qwerty';"
```

---

## Create Database and 
```shell
psql -c 'create database bookings;'
psql -c 'grant all privileges on database bookings to bookings_service;'
```

---

## Drop Database

```shell
psql -c "drop database bookings_service;"
```