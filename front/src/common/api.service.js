import axios from "axios"
import API_URL from "@/common/config"

function token() {
  return localStorage.getItem("token")
}

const ApiService = {
  init() {
    axios.defaults.baseURL = API_URL
    this.setHeader()
  },

  setHeader() {
    axios.defaults.headers.common["Rabbit-Fur"] = token()
    axios.defaults.headers.common["Content-Type"] = "application/json;charset=utf8"
  },

  get(resource, slug="") {
    return axios.get(`${resource}/${slug}`)
  },

  post(resource, data) {
    return axios.post(`${resource}`, data)
  },

  patch(resource, slug, data) {
    return axios.patch(`${resource}/${slug}`, data)
  },

  delete(resource, slug) {
    return axios.delete(`${resource}/${slug}`)
  }
}

export default ApiService