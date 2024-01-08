# thoughtspace

## Schema

Simple schema

```md
Users

- id
- username
- email

Posts

- id
- title
- description
- authorId -> Users.id

Comments

- id
- postId -> Posts.id
- parentId (nullable)
- comment
- attachmentUrl (nullable)
- authorId -> Users.id
```
