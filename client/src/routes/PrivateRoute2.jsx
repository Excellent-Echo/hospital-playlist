import React from 'react'

import { Redirect } from 'react-router-dom'
import { useSelector } from 'react-redux'

const PrivateRoute = () => {
	const userProfileData = useSelector((state) => state.userProfile)
	console.log(userProfileData)

	// return userProfileData.user_id === '' ? <Redirect to="/login" /> : <span>ini private route</span>
	return userProfileData.user_id === 'xx' ? <Redirect to="/login" /> : <Redirect to="/profile" />
}

export default PrivateRoute
