import axios from "axios"
import { useContext, useState } from "react"
import { MyContextAuth } from "./Auth"
import ModalComponent from "./Modal"
import { toast } from "react-toastify"

export default function CreateProduct(props) {

    const [description, setDescription] = useState('')
    const [price, setPrice] = useState(0.0)

    const { getToken } = useContext(MyContextAuth)

    function createProduct() {
        const token = getToken()
        const config = { headers: { Authorization: "Bearer " + token } }

        axios.post("http://localhost:8080/creator/product", { description, price }, config)
            .then((resposne) => {
                toast.success('Produto cadastrado!', {
                    position: "top-right",
                    autoClose: 5000,
                    hideProgressBar: false,
                    closeOnClick: true,
                    pauseOnHover: true,
                    draggable: true,
                    progress: undefined,
                    theme: "light",
                });
                props.show(false)
            })
            .catch((error) => {
                const errorMessage = error.response.data.Info.Message
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
    }


    return (
        <ModalComponent>
            <div className="d-flex justify-content-lg-between align-items-center " style={{ marginBottom: '10px' }}>
                <div>Cadastrar Produto</div>
                <div>
                    <button type="button" className="btn-close" aria-label="Close" onClick={() => props.show(false)}></button>
                </div>
            </div>
            <div className="input-group mb-3">
                <span className="input-group-text" id="basic-addon1">Descricao</span>
                <input type="text" className="form-control" aria-label="descricao" aria-describedby="basic-addon1" onChange={e => setDescription(e.target.value)} />
            </div>

            <div className="input-group mb-3">
                <span className="input-group-text">$</span>
                <input type="text" className="form-control" aria-label="Amount (to the nearest dollar)" onChange={e => setPrice(e.target.value)} />
                <span className="input-group-text">.00</span>
            </div>

            <div>
                <button type="button" className="btn btn-dark" onClick={() => createProduct()}>Adicionar</button>
            </div>


        </ModalComponent>
    )


}