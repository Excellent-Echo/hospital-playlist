import axios from 'axios'

const HelperAPI = axios.create({
	baseURL: 'http://localhost:8080'
})

export default HelperAPI
