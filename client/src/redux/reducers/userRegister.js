import {
	USER_REGISTER_RESET_FORM,
	USER_REGISTER_SET_FULLNAME,
	USER_REGISTER_SET_BIRTHDATE,
	USER_REGISTER_SET_ADDRESS,
	USER_REGISTER_SET_GENDER,
	USER_REGISTER_SET_EMAIL,
	USER_REGISTER_SET_PASSWORD,
	USER_REGISTER_SET_PASSWORD,
	USER_REGISTER_SET_ERROR_MESSAGE,
	USER_REGISTER_SET_SUCCESS_MESSAGE,
	USER_REGISTER_START_LOADING,
	USER_REGISTER_STOP_LOADING
} from '../actionTypes/userRegister'

const initialState = {
	firstName: '',
	lastName: '',
	email: '',
	password: '',
	fullName: '',
	birthDate: '',
	address: '',
	gender: '',
	errorMessage: '',
	successMessage: '',
	isLoading: false
}

const userRegisterReducer = (state = initialState, action) => {
	switch (action.type) {
		case USER_REGISTER_RESET_FORM:
			return {
				...initialState
			}
		case USER_REGISTER_SET_EMAIL:
			return {
				...state,
				email: action.payload.email
			}
		case USER_REGISTER_SET_PASSWORD:
			return {
				...state,
				password: action.payload.password
			}
		case USER_REGISTER_SET_FULLNAME:
			return {
				...state,
				firstName: action.payload.fullName
			}
		case USER_REGISTER_SET_BIRTHDATE:
			return {
				...state,
				firstName: action.payload.birthDate
			}
		case USER_REGISTER_SET_ADDRESS:
			return {
				...state,
				firstName: action.payload.address
			}
		case USER_REGISTER_SET_GENDER:
			return {
				...state,
				firstName: action.payload.gender
			}
		case USER_REGISTER_SET_ERROR_MESSAGE:
			return {
				...state,
				errorMessage: action.payload.errorMessage
			}

		case USER_REGISTER_SET_SUCCESS_MESSAGE:
			return {
				...state,
				successMessage: action.payload.successMessage
			}

		case USER_REGISTER_START_LOADING:
			return {
				...state,
				isLoading: true
			}

		case USER_REGISTER_STOP_LOADING:
			return {
				...state,
				isLoading: false
			}

		default:
			return state
	}
}
