import userRegisterReducer from './userRegister'
import userProfileReducer from './userProfile'
import userLoginReducer from './userLogin'

const rootReducer = {
	userRegister: userRegisterReducer,
	userProfile: userProfileReducer,
	userLogin: userLoginReducer
}

export default rootReducer
