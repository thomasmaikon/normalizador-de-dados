import axios from "axios"
import { useContext, useEffect, useState } from "react"
import { MyContextAuth } from "./Auth"

export default function Historical() {
    const { getToken } = useContext(MyContextAuth)

    const [transactionsHistorical, setTransactionsHistorical] = useState()

    useEffect(() => {
        const token = getToken()
        const config = { headers: { Authorization: "Bearer " + token } };
        axios.get('http://localhost:8080/creator/historical', config)
            .then((response) => {
                const listTransactionsHistorical = response.data.Info
                setTransactionsHistorical(listTransactionsHistorical)
            })
            .catch((error)=>{
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
    }, [])

    function showTransactions() {

        return transactionsHistorical.map((data, index) => {
            const date = new Date(data.Date);
            const day = String(date.getDate()+1).padStart(2, "0");
            const month = String(date.getMonth() + 1).padStart(2, "0");
            const year = date.getFullYear();
            const formattedDate = `${day}/${month}/${year}`;
            const linkAfiliate = "/afiliate/" + data.AfiliateId
            
            return (
                <tr key={index}>
                    <td><a href={linkAfiliate} >{data.AfiliateName}</a></td>
                    <td>{data.ProductDescription}</td>
                    <td>{data.TransactionDescription}</td>
                    <td>{data.Value}</td>
                    <td>{formattedDate}</td>
                </tr>
            )
        })


    }

    if (!transactionsHistorical) {
        return
    }

    return (
        <div className="table table-striped table-hover" style={{ overflow: 'auto', maxHeight:'500px' }}>
            <table className="table table-striped table-hover" >
                <thead>
                    <tr>
                        <th scope="col">Afiliado</th>
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

    )
}