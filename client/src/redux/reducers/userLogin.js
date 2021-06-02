import {
	USER_LOGIN_SET_EMAIL,
	USER_LOGIN_SET_PASSWORD,
	USER_LOGIN_SET_ERROR_MESSAGE,
	USER_LOGIN_START_LOADING,
	USER_LOGIN_STOP_LOADING,
	USER_LOGIN,
	USER_LOGOUT
} from '../actionTypes/userLogin'

const initialState = {
	email: '',
	password: '',
	errorMessage: '',
	isLoading: false,
	isAuth: false
}

const userLogin = (state = initialState, action) => {
	switch (action.type) {
		case USER_LOGIN_SET_EMAIL:
			return {
				...state,
				email: action.payload.email
			}
		case USER_LOGIN_SET_PASSWORD:
			return {
				...state,
				password: action.payload.password
			}
		case USER_LOGIN_SET_ERROR_MESSAGE:
			return {
				...state,
				errorMessage: action.payload.errorMessage
			}

		case USER_LOGIN_START_LOADING:
			return {
				...state,
				isLoading: true
			}

		case USER_LOGIN_STOP_LOADING:
			return {
				...state,
				isLoading: false
			}
		case USER_LOGIN:
			return {
				...state,
				isAuth: true
			}
		case USER_LOGOUT:
			return {
				...state,
				isAuth: false
			}

		default:
			return state
	}
}

export default userLogin
