# Computer Alchemy API

A RESTful API for the Computer Alchemy game that manages elements and their combinations. Built with Go and Neo4j.

## 🚀 Features

- Element management (create, retrieve)
- Element combination system
- Swagger documentation
- Neo4j graph database integration

## 📋 Prerequisites

Before running this project, make sure you have the following installed:
- Go 1.x
- Neo4j Database
- Git

## ⚙️ Environment Variables

Create a `.env` file in the root directory with the following variables:

```
PORT=
NEO4J_URI=
NEO4J_USER=
NEO4J_PASSWORD=
```

## 🛠️ Running the project locally

1. Clone the repository
```
git clone https://github.com/your-username/computer-alchemy-api.git
```

2. Install dependencies
```
go mod download
```

3. Run the server
```
go run main.go
```

4. Access the Swagger UI
Open your browser and navigate to `http://localhost:port/docs/index.html` to interact with the API.

## 📚 API Documentation

### Available Endpoints

#### Elements
- `GET /api/elements` - Get all elements
- `GET /api/elements/{name}` - Get element by name

#### Combinations
- `GET /api/combinations` - Get all combinations
- `POST /api/combinations` - Combine two elements
- `GET /api/combinations?element={elementName}` - Get combinations for specific element
- `GET /api/combinations/result?resultingElement={elementName}` - Get combinations that result in specific element

For detailed information about the API, including endpoints, request/response formats, and error codes, please refer to the Swagger UI `/docs/index.html`.

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a pull request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📜 License

This project is open-sourced under the MIT License - see the LICENSE file for details.

## 👥 Authors

- **PandaX185** - *Initial work*

## 🙏 Acknowledgments

- Neo4j Go Driver team
- Gorilla Mux contributors
- Swagger team
