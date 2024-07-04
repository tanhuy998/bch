import ErrorResponse from "../backend/error/errorResponse";
import useNotFoundRedirection from "./useNotFoundRedirection";

/**
 * 
 * @param {ErrorResponse} error 
 * @returns 
 */
export function useEndpointResponseErrorHandler() {

    const redirectNotFound = useNotFoundRedirection();

    return (error) => {

        // switch (error.responseObject.status) {

        //     case 404:
        //         redirectNotFound();
        //         return;
        // }

        const status = error.responseObject.status;

        if (status >= 400 && status < 500) {

            redirectNotFound();
        }

        if (status >= 500) {

            
        }
    }
}