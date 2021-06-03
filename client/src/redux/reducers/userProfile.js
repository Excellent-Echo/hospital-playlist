import { USER_PROFILE_FETCH_DATA } from '../actionTypes/userProfile'

const initialState = {
	id: '',
	no_handphone: '',
	birth_date: '',
	gender: '',
	address: '',
	user_id: '',
	created_at: '',
	updated_at: ''
}

const userProfileReducer = (state = initialState, action) => {
	switch (action.type) {
		case USER_PROFILE_FETCH_DATA:
			return {
				...state,
				...action.payload.user
			}

		default:
			return state
	}
}

export default userProfileReducer
