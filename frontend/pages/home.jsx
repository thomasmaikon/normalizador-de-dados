import { useContext, useEffect, useState } from "react"
import Creator from "../components/Creator"
import CreateAfiliated from "../components/CreateAfiliated"
import CreateProduct from "../components/CreateProduct"
import UploadFile from "../components/UploadFileForm"
import Historical from "../components/Historical"
import axios from "axios"
import { MyContextAuth } from "../components/Auth"
import { useRouter } from "next/router"
import { toast } from "react-toastify"
import LogOut from "../components/Logout"

export default function CreatorURL() {

    const { getToken } = useContext(MyContextAuth)
    const router = useRouter()
    const [oppenModalCreator, setOpenModalCreator] = useState(false)
    const [openModalAfiliate, setOpenModalAfiliate] = useState(false)
    const [oppenModalProduct, setOpenModalProduct] = useState(false)
    const [creatorName, setCreatorName] = useState(undefined)
    const [creatorAmount, setCreatorAmount] = useState(0.0)


    useEffect(() => {
        const isAuth = getToken() != undefined
        if (isAuth) {
            const token = getToken()
            const config = { headers: { Authorization: "Bearer " + token } };
            axios.get('http://localhost:8080/creator', config)
                .then((response) => {
                    const name = response.data.Info.Name
                    const amount = response.data.Amount
                    setCreatorName(name)
                    setCreatorAmount(amount)
                })
                .catch((error) => {
                    const errorMessage = error.response.data.Info.Message
                    setCreatorName("")
                    toast.error(errorMessage, {
                        position: "top-right",
                        autoClose: 5000,
                        hideProgressBar: false,
                        closeOnClick: true,
                        pauseOnHover: true,
                        draggable: true,
                        progress: undefined,
                        theme: "light",
                    });
                })
        } else {
            router.push('/login')
        }
    },[])

    if(creatorName == undefined){
        return 
    }

    if (creatorName == "") {
        return (
            <div className="container text-center d-flex flex-column justify-content-center align-items-center " style={{ height: 'inherit' }}>
                Voce nao tem creator cadastrado, cadastre um e seus respectivos afiliados e produtos
                <div className="d-flex" style={{ minWidth: '300px', justifyContent: 'space-between', padding: '15px' }}>
                    <div>
                        <button type="button" className="btn btn-dark" onClick={e => setOpenModalCreator(true)}>
                            Novo creator
                        </button>
                    </div>
                </div>
                {oppenModalCreator && <Creator show={setOpenModalCreator} />}
            </div>
        )
    } else {
        return (
            <div className="container text-center d-flex flex-column justify-content-center align-items-center " 
            style={{ height: '80%', boxShadow:' rgba(50, 50, 93, 0.25) 0px 30px 60px -12px, rgba(0, 0, 0, 0.3) 0px 18px 36px -18px' }}>
                <LogOut ></LogOut>
                <div>
                   <h5><b>{creatorName}</b></h5> <h5>{creatorAmount} R$</h5>
                </div>

                <div className="d-flex" style={{ minWidth: '300px', justifyContent: 'space-between', padding: '15px' }}>
                    {creatorName && <div><button type="button" className="btn btn-dark" onClick={e => setOpenModalAfiliate(true)}>Novo afiliado</button></div>}
                    {creatorName && <div><button type="button" className="btn btn-dark" onClick={e => setOpenModalProduct(true)}>Novo produto</button></div>}
                </div>

                <UploadFile />
                <Historical />

                {openModalAfiliate && <CreateAfiliated show={setOpenModalAfiliate} />}
                {oppenModalProduct && <CreateProduct show={setOpenModalProduct} />}
            </div>
        )
    }
}