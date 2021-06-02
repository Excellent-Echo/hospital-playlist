import HospitalAPI from '../../api/HospitalAPI'

import {
	USER_LOGIN_SET_EMAIL,
	USER_LOGIN_SET_PASSWORD,
	USER_LOGIN_SET_ERROR_MESSAGE,
	USER_LOGIN_START_LOADING,
	USER_LOGIN_STOP_LOADING,
	USER_LOGIN,
	USER_LOGOUT
} from '../actionTypes/userLogin'

import userProfileAction from '../actions/userProfile'

const setEmail = (email) => {
	return {
		type: USER_LOGIN_SET_EMAIL,
		payload: {
			email: email
		}
	}
}

const setPassword = (password) => ({
	type: USER_LOGIN_SET_PASSWORD,
	payload: {
		password: password
	}
})

const setErrorMessage = (errorMessage) => ({
	type: USER_LOGIN_SET_ERROR_MESSAGE,
	payload: {
		errorMessage: errorMessage
	}
})

const startLoading = () => ({
	type: USER_LOGIN_START_LOADING
})

const stopLoading = () => ({
	type: USER_LOGIN_STOP_LOADING
})

const userLoginAuth = () => ({
	type: USER_LOGIN
})

const userLogoutAuth = () => ({
	type: USER_LOGOUT
})

const login = (email, password, history) => async (dispatch) => {
	try {
		dispatch(startLoading())
		dispatch(setErrorMessage(''))

		const loginData = {
			email: email,
			password: password
		}

		const user = await HospitalAPI({
			method: 'POST',
			url: '/users/login',
			data: loginData
		})

		const accessToken = user.data.data.token

		localStorage.setItem('accessToken', accessToken)
		localStorage.setItem('isAuth', true)

		// const userProfile = await HospitalAPI({
		// 	method: 'GET',
		// 	url: '/users',
		// 	headers: {
		// 		Authorization: user.data.data.token
		// 	}
		// })

		// dispatch(userProfileAction.setProfileData(userProfile.data.data))
		dispatch(userLoginAuth())
		history.push('/admin')
		dispatch(stopLoading())
	} catch (error) {
		console.log(error)
		// dispatch(setErrorMessage(error.response.data.data.errors || ['internal server error']))
		dispatch(stopLoading())
	}
}

const userLogin = {
	setEmail,
	setPassword,
	login
}

export default userLogin
