class ApiClient {
  // let baseURL = "http://localhost:3000/api";
  // let token = null;
  constructor(baseURL) {
    this.baseURL = baseURL;
  }

  // Fonction pour le login
  async login(email, password) {
    return this._post("/login", { email, pass: password });
  }

  // Méthodes génériques pour les requêtes CRUD
  async create(entity, data) {
    return this._post(`/${entity}`, data);
  }

  async read(entity, id) {
    return this._get(`/${entity}/${id}`);
  }

  async update(entity, id, data) {
    return this._put(`/${entity}/${id}`, data);
  }

  async delete(entity, id) {
    return this._delete(`/${entity}/${id}`);
  }

  // Méthodes internes pour simplifier les appels à `fetch`
  async _get(url) {
    const response = await fetch(`${this.baseURL}${url}`, {
      method: "GET",
      headers: this._getHeaders(),
    });
    return this._handleResponse(response);
  }

  async _post(url, data) {
    const response = await fetch(`${this.baseURL}${url}`, {
      method: "POST",
      headers: this._getHeaders(),
      body: JSON.stringify(data),
    });
    return this._handleResponse(response);
  }

  async _put(url, data) {
    const response = await fetch(`${this.baseURL}${url}`, {
      method: "PUT",
      headers: this._getHeaders(),
      body: JSON.stringify(data),
    });
    return this._handleResponse(response);
  }

  async _delete(url) {
    const response = await fetch(`${this.baseURL}${url}`, {
      method: "DELETE",
      headers: this._getHeaders(),
    });
    return this._handleResponse(response);
  }

  // Headers par défaut, notamment pour les JSON
  _getHeaders() {
    return {
      "Content-Type": "application/json",
      Authorization: "Bearer YOUR_TOKEN", // Remplacer par le token de session après login
    };
  }

  // Gestion de la réponse pour capturer les erreurs et obtenir le JSON
  async _handleResponse(response) {
    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }
    return response.json();
  }
}

// // Utilisation
// const api = new ApiClient('https://api.example.com');

// // Exemple de login
// api.login('mon@email.com', '123pass')
//    .then(data => console.log('Login successful:', data))
//    .catch(error => console.error('Login error:', error));

// // Exemple CRUD pour l'entité "user"
// api.create('user', { name: 'John Doe', email: 'john@example.com' })
//    .then(data => console.log('User created:', data))
//    .catch(error => console.error('Error creating user:', error));

// api.read('user', 'USER_ID')
//    .then(data => console.log('User data:', data))
//    .catch(error => console.error('Error fetching user:', error));

// api.update('user', 'USER_ID', { email: 'new@example.com' })
//    .then(data => console.log('User updated:', data))
//    .catch(error => console.error('Error updating user:', error));

// api.delete('user', 'USER_ID')
//    .then(data => console.log('User deleted:', data))
//    .catch(error => console.error('Error deleting user:', error));
