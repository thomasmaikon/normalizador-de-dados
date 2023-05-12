import { useContext, useEffect, useState } from "react"
import { MyContextAuth } from "./Auth"
import axios from "axios"
import { toast } from "react-toastify"
import { useRouter } from "next/router"

export default function AfiliateHistorical(props) {
    const { getToken } = useContext(MyContextAuth)
    const router = useRouter()
    const [transactionsHistorical, setTransactionsHistorical] = useState([])
    const [afiliateName, setAfiliateName] = useState(0.0)
    const [amountValue, setAmountValue] = useState(0.0)

    useEffect(() => {
        const isAuth = getToken() != undefined
        if (isAuth) {
            const token = getToken()
            const config = { headers: { Authorization: "Bearer " + token } };
            axios.get('http://localhost:8080/creator/historical/afiliate/' + props.afiliate, config)
                .then((response) => {
                    const listTransactionsHistorical = response.data.Info.AfiliateHistoricals
                    const name = response.data.Info.AfiliateHistoricals[0].AfiliateName
                    const amount = response.data.Info.Amount
                    setTransactionsHistorical(listTransactionsHistorical)
                    setAfiliateName(name)
                    setAmountValue(amount)
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
        } else {
            router.push('/login')
        }
    }, [])

    function showTransactions() {

        return transactionsHistorical.map((data, index) => {
            const date = new Date(data.Date);
            const day = String(date.getDate()).padStart(2, "0");
            const month = String(date.getMonth() + 1).padStart(2, "0");
            const year = date.getFullYear();
            const formattedDate = `${day}/${month}/${year}`;

            return (
                <tr key={index}>
                    <td>{data.ProductDescription}</td>
                    <td>{data.TransactionDescription}</td>
                    <td>{data.Value}</td>
                    <td>{formattedDate}</td>
                </tr>
            )
        })
    }

    return (
        <div className="container text-center d-flex flex-column justify-content-center align-items-center " style={{ height: 'inherit' }}>
            <div style={{marginBottom:'30px'}}>
                <h5>{afiliateName} : {amountValue}</h5>
            </div>
            <div className="table table-striped table-hover" style={{ overflow: 'auto', maxHeight: '500px' }}>
                <table className="table table-striped table-hover">
                    <thead>
                        <tr>
                            <th scope="col">Produto</th>
                            <th scope="col">Transacao</th>
                            <th scope="col">Valor</th>
                            <th scope="col">Data</th>
                        </tr>
                    </thead>
                    <tbody>
                        {showTransactions()}
                    </tbody>
                </table>
            </div>
        </div>
    )

}