
# CLI store backend

Backend written in Go Fiber using sqlite database for CLI store management tool. 


## API Reference

#### Get all items

```bash
  GET /api
```
#### Get item

```bash
  GET /api/items/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

#### New item
```bash
  POST /api/item
```
#### Sell (update) item

```bash
    PUT /api/item/sell/${id}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to sell |


#### Delete item

```bash
    DELETE /api/item/${id}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to delete |
## Installation

Install backend using following after cloning this repository: 

```bash 
  go mod tidy
```
Run the server using :


```bash 
  air server --port 8000
```
## Authors

- [@dibyajyoti](https://www.github.com/dibyajyoti-mandal)

