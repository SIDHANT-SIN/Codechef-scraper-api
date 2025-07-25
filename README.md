
<div>

## About 

The **CodeStats API** is a lightweight RESTful API designed to scrape user statistics and contest history from CodeChef profiles. Built with Go and the Gin framework, it provides fast, structured JSON responses for developers and competitive programming enthusiasts.

</div>

## API URL 🌐

```bash
https://codechef-scraper-v1-2.onrender.com/
```

## ‼️ Note

For development, run the API locally to avoid rate limits.

---

### Local Deploy 🛠️

To run the CodeStats API locally for development or testing:

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/SIDHANT-SIN/codechef-scraper-api.git
   cd codechef-scraper-api
   ```

2. **Set Up Environment Variables**:
   - Create an `.env` file in the project root:
     ```bash
     PORT=8080 
     RATE_LIMIT_REQUESTS=YOUR_LIMIT
     RATE_LIMIT_WINDOW_HOURS=YOUR_WINDOW
     ```


3. **Run with Go**:
   - Ensure Go 1.24 or above is installed.
   - Install dependencies:
     ```bash
     go mod tidy
     ```
   - Run the API:
     ```bash
     go run main.go
     ```
   - Access at `http://localhost:8080`.
  
---


   

### Dockerized Setup 🐳

This section provides instructions for running the API using Docker. The Docker image is publicly available on [Docker Hub](https://hub.docker.com/r/sidhant3singh/codechef-scraper).

1.  **Run with Pre-built Docker Image (from Docker Hub)**:
    For the easiest setup, pull the pre-built Docker image.
    
    ```bash
    docker pull sidhant3singh/codechef-scraper:v1.2
    ```
    
    Then, run the container with environment variables:
    
    ```bash
    docker run -p 8080:8080 -e PORT=8080 -e RATE_LIMIT_REQUESTS=YOUR_LIMIT -e RATE_LIMIT_WINDOW_HOURS=YOUR_WINDOW sidhant3singh/codechef-scraper:v1.2
    ```
    
    This will start the API, accessible at `http://localhost:8080`.

---


## Endpoints 🚀

| Endpoint | Description | Demo |
| :------- | :---------- | :--- |
| <code>GET /user/:username</code> | Retrieve detailed profile stats for a CodeChef user | <a href="https://codechef-scraper-v1-2.onrender.com/user/tourist" target="_blank" >click here</a> |
| <code>GET /contests/:username</code> | Fetch contest history for a user | <a href="https://codechef-scraper-v1-2.onrender.com/contests/tourist" target="_blank">click here</a> |

## For user stats - 

### Example Request 

```bash
curl https://codechef-scraper-v1-2.onrender.com/user/your_codechef_username
```

### Example Response

```json
{
  "username": "your_username",
  "name": "your_name",
  "stars": "★★★",
  "rating": "1600",
  "highest_rating": "1700",
  "global_rank": "1620",
  "country_rank": "1320",
  "country": "India",
  "institution": "University"
}
```

## For contests participated - 

### Example Request 

```bash
curl https://codechef-scraper-v1-2.onrender.com/contests/your_codechef_username
```

### Example Response

```json
[
  {
    "code": "FEB23",
    "getyear": "2023",
    "getmonth": "2",
    "getday": "10",
    "rating": "1750",
    "rank": "200",
    "name": "February Challenge 2023",
    "end_date": "2023-02-10 18:00:00"
  },
  {
    "code": "MAR23",
    "getyear": "2023",
    "getmonth": "3",
    "getday": "12",
    "rating": "1780",
    "rank": "180",
    "name": "March Cook-Off 2023",
    "end_date": "2023-03-12 20:00:00"
  }
]
```

### 💡 Rate Limit

The API enforces a rate limit of 10 requests per hour per client IP to ensure fair usage and safety from Denial of Service (DoS) attack.

