import HospitalAPI from '../../api/HospitalAPI'

import { USER_PROFILE_FETCH_DATA } from '../actionTypes/userProfile'

const fetchProfileData = () => async (dispatch) => {
	try {
		const user = await HospitalAPI({
			method: 'GET',
			headers: {
				Authorization: localStorage.getItem('accessToken')
			}
		})

		dispatch({
			type: USER_PROFILE_FETCH_DATA,
			payload: {
				user: user.data.data
			}
		})
	} catch (error) {
		console.log(error)
	}
}

const setProfileData = (userProfile) => {
	return {
		type: USER_PROFILE_FETCH_DATA,
		payload: {
			user: userProfile
		}
	}
}

const userProfile = {
	fetchProfileData,
	setProfileData
}

export default userProfile
