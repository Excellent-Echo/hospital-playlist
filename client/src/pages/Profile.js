import React from 'react'
import { useSelector } from 'react-redux'

const ProfilePage = () => {
	const userProfileData = useSelector((state) => state.userProfile)
	return (
		<>
			<span>Profile Pages</span>
			<pre>{JSON.stringify(userProfileData)}</pre>
		</>
	)
}

export default ProfilePage
