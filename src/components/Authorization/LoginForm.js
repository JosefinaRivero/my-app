import  React from "react"
import CONFIG from "../../app.config.json"


export const LoginForm = ({ setIsAuthenticated, setUser }) => {
    const [email, setEmail] = React.useState("");
    const [pwd, setPwd] = React.useState("");

    const login = async () => {
        try {
            const response = await fetch(CONFIG.BASEURL + "/login", {
                method: "POST", // *GET, POST, PUT, DELETE, etc.
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({email: email, pwd: pwd}),
            });
            response.json().then(result => {
                setIsAuthenticated(true)
                setUser(result)
            });
        } catch (error) {

        }
    }

    return (<div>
        <p>Login:</p>
        <label>Email</label>
        <input value={email} onChange={(e) => { setEmail(e.target.value) }}/>
        <label>Pwd</label>
        <input value={pwd} type="password" onChange={(e) => { setPwd(e.target.value) }}/>
        <button onClick={() => { login() }}>Login</button>
    </div>

    );
};