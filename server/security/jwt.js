import jwt from 'jsonwebtoken'

const secretKey = process.env.SECRETKEY || 'secret'.repeat(3)

function login(req, res) {
    let user = req.body
    // ensure that token registered to user in db is invalid
        // check if user is even in db
        // check if password has is valid
        // check if user has a token
        // check if user's token is valid
    if(userInDB){
        if (userHasToken){
            res.status(400).send({
                message: ''
            })
        } || userTokenInvalid)){ 
            // create token
            let newToken = jwt.sign({user:user,iat:Math.floor(Date.now() / 1000) - 30},secretKey,{expiresIn:60*10,subject:user.name})
            // store new token in db
            res.status(200).send({
                token:newToken
            })
}else{
        res.status(400).send({
            message: ''
        })
    }
}

function verifyJWT(req, res)