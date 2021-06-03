import React from 'react'
import { Route, Redirect } from 'react-router-dom'

const PrivateRoute = ({ component: Component, role, auth, ...rest }) => {
	return (
		<Route
			{...rest}
			render={(props) =>
				auth === true && role === 'admin' ? <Component {...props} /> : <Redirect to="/login" />
			}
		/>
	)
}

export default PrivateRoute
