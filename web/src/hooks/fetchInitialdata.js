import { useEffect, useState } from "react";
import HttpEndpoint from "../backend/endpoint";

/**
 * 
 * @param {*} endpoint 
 * @param {*} args 
 */
export default function fetchInitialData(endpoint, args = []) {

    if (!(endpoint instanceof HttpEndpoint)) {

        throw new Error('invalid endpoint object ot fetch data');
    }

    const [data, setData] = useState(null);

    useEffect(() => {

        
    })

    return [data, setData];
}