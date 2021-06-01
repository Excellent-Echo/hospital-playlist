import HospitalAPI from '../../api/HospitalAPI'

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

const resetForm = () => {
	return {
		type: USER_REGISTER_RESET_FORM
	}
}

const setEmail = (email) => {
	return {
		type: USER_REGISTER_SET_EMAIL,
		payload: {
			email: email
		}
	}
}

const setPassword = (password) => {
	return {
		type: USER_REGISTER_SET_PASSWORD,
		payload: {
			password: password
		}
	}
}

const setFullName = (fullName) => {
	return {
		type: USER_REGISTER_SET_FULLNAME,
		payload: {
			fullName: fullName
		}
	}
}

const setBirthDate = (birthDate) => {
	return {
		type: USER_REGISTER_SET_BIRTHDATE,
		payload: {
			birthDate: birthDate
		}
	}
}

const setAddress = (address) => {
	return {
		type: USER_REGISTER_SET_ADDRESS,
		payload: {
			address: address
		}
	}
}

const setGender = (gender) => {
	return {
		type: USER_REGISTER_SET_GENDER,
		payload: {
			gender: gender
		}
	}
}

const setErrorMessage = (errorMessage) => {
	return {
		type: USER_REGISTER_SET_ERROR_MESSAGE,
		payload: {
			errorMessage: errorMessage
		}
	}
}

const setSuccessMessage = (successMessage) => {
	return {
		type: USER_REGISTER_SET_SUCCESS_MESSAGE,
		payload: {
			successMessage: successMessage
		}
	}
}

const startLoading = () => {
	return {
		type: USER_REGISTER_START_LOADING
	}
}

const stopLoading = () => {
	return {
		type: USER_REGISTER_STOP_LOADING
	}
}

const register = (email, password, fullName, birthDate, address, gender) => async (dispatch) => {
	try {
		dispatch(startLoading())
		dispatch(setSuccessMessage(''))
		dispatch(setErrorMessage(''))
		const submitData = {
			email: email,
			password: password,
			fullName: fullName,
			birthDate: birthDate,
			address: address,
			gender: gender
		}

		const user = await HospitalAPI({
			method: 'POST',
			url: '/users/register',
			data: submitData
		})

		dispatch(setSuccessMessage('Registrasi Sukses, Silahkan login'))
		dispatch(stopLoading())
	} catch (error) {
		console.log(error.response)
		dispatch(setErrorMessage(error.response.data.data.errors || ['internal server error']))
		dispatch(stopLoading())
	}
}

const userRegister = {
	resetForm,
	setEmail,
	setPassword,
	setFullName,
	setBirthDate,
	setAddress,
	setGender,
	register
}

export default userRegister
