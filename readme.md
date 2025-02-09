<!--suppress ALL -->
<p align="center">
    <img src="https://github.com/ShiftSad/Shorty/blob/master/upload/header.png" alt="Shorty" />
    <br/>
    <b>
        <a href="https://shorty.laranjaazul.online">Try Now!</a>
    </b> —
    <b>
        <a href="http://github.com/ShiftSad/Shorty/issues">Issues</a>
    </b>
</p>
<br/>

## ⚡ Introduction

**Shorty** is a free, open-source URL shortening service. Designed to be blazingly fast™.

## 🔥 Tech Stack

![Go](https://img.shields.io/badge/Go-1B73BA?style=for-the-badge&logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-00A99D?style=for-the-badge&logo=go&logoColor=white)
![Gorm](https://img.shields.io/badge/Gorm-764ABC?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![HTML5](https://img.shields.io/badge/HTML5-E34F26?style=for-the-badge&logo=html5&logoColor=white)
![CSS3](https://img.shields.io/badge/CSS3-1572B6?style=for-the-badge&logo=css3&logoColor=white)
![JavaScript](https://img.shields.io/badge/JavaScript-F7DF1E?style=for-the-badge&logo=javascript&logoColor=black)

## 🚀 How to Self-Host

You can deploy this project on **Railway** with just one click:  

[![Deploy on Railway](https://railway.com/button.svg)](https://railway.com/template/ie90_N?referralCode=H3NRbK)

### Manual Setup

1. **Clone the Repository**  
   ```sh
   git clone https://github.com/ShiftSad/Shorty.git
   cd Shorty
   ```

2. **Set Up PostgreSQL**  
   Ensure you have a PostgreSQL database running. Then, set the **DSN** (Data Source Name) environment variable with your PostgreSQL connection URL.

   **On Linux/macOS:**
   ```sh
   export DSN="postgresql://postgres:pa$$word@localhost:5432/postgres"
   ```

   **On Windows (PowerShell):**
   ```powershell
   $env:DSN="postgresql://postgres:pa$$word@localhost:5432/postgres"
   ```

3. **Run the Application**  
   ```sh
   go run main.go
   ```

That's it! Your application should now be running locally. 🚀  

## 👀 Visuals

![Visuals](https://github.com/ShiftSad/Shorty/blob/master/upload/visual.png)
