import { pillTabStyle } from "../../contexts/tab.context";


const customTabContextvalue = {
    ...pillTabStyle,
    li: {
        ...pillTabStyle.li,
        style: {
            display: 'inline-block',
        }
    },
    ul: {
        ...pillTabStyle.ul,
        style: {
            backGround: 'rgba(255, 255, 255, 0)',

        }
    }
}

export default customTabContextvalue;