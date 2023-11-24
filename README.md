# WSCHAT Server
**Websocket chat app server, build with golang**

This project is part of Talenthub batch 11 "Scalable Web Service With Go" submission

---

REST API ROUTES:
GET `/session` >> return all active/avaliable session chat
GET `/session/create?type=<PRIVATE/GROUP>` >> create new session, return Session
POST `/user/create` body :` { name: "yourname" }` >> create new user


WEBSOCKET
GET `/ws?session_id=<SESSION_ID>&user_id=<USER_ID>`


&copy; Ahmad Ma'ruf - 2023