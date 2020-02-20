# Notes

## api

### api设计
1. 用户
- 创建(注册)用户
    - URL: /user
    - Method : POST
    - SC: 201(created), 400(bad request), 500(internal error)

- 用户登录
    - URL: /user/:username
    - Method: POST
    - SC: 200(OK), 400, 500

- 获取用户登陆信息
    - URL: /user/:username
    - Method: GET
    - SC: 200, 400, 401(并没有验证), 403(验证了，但没有对应权限), 500

- 用户注销
    - URL: /user/:username
    - Method: DELETE
    - SC: 204(no content), 400, 401, 403, 500

2. 视频

- List all videos:
    - URL: /user/:username/videos
    - Method: GET
    - SC: 200, 400, 500

- Get one video:
    - URL: /user/:username/videos/:vid-id
    - Method: GET
    - SC: 200, 400, 500

- Delete one video
    - URL: /user/:username/videos/:vid-id
    - Method: DELETE
    - SC: 204, 400, 401, 403, 500

3. 评论

- Show comments:
    - URL: /videos/:vid-id/comments
    - Method: GET
    - SC: 200, 400, 500

- Post one comment
    - URL: /videos/:vid-id/comments
    - Method: POST
    - SC: 201, 400, 500 

- Delete one comment
    - URL: /videos/:vid-id/comments/:comment-id
    - Method: DELETE
    - SC: 204, 400, 401, 403, 500


source createTable.sql;

go get -u github.com/julienschmidt/httprouter
go get -u github.com/go-sql-driver/mysql