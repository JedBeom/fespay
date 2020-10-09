export default function (to, from, next) {
    let token = localStorage.getItem("token")
    if (!token) {
        next({name: "login"})
        return
    }
    next()
}