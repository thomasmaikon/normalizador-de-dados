import { useRouter } from "next/router";
import AfiliateHistorical from "../../components/AfiliateHistorical";

export default function Afiliate(){
    const router = useRouter();
    const { id } = router.query;
    
    if (!id){
      return 
    }

    return (
      <AfiliateHistorical afiliate={id}/>
    );
}