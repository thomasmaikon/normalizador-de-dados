import { useContext } from "react"
import { MyContextAuth } from "./Auth"
import Cookies from "js-cookie"
import { useRouter } from "next/router"

export default function LogOut() {

    const router = useRouter()

    function logout() {
        Cookies.remove('hubla-token')
        router.push('/login')
    }

    return (
        <div style={{position:'fixed', top:'20px', left: '20px'}}>
            <button type="button" className="btn btn-outline-danger" onClick={()=>logout()}>Sair</button>
        </div>
    )
}