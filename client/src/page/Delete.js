import { Col, Container , Row } from "react-bootstrap";
import { useNavigate, useParams } from "react-router-dom";
import axios from "axios"

const Delete = () => {

    const params = useParams()
    const navvigate = useNavigate()
    const handleDelete = async () => {
        try {
            const apiUrl = process.env.REACT_APP_API_ROOT + "/" + params.id;
            const response = await axios.delete(apiUrl);
            if(response.status === 200) {
                navvigate("/", {
                    state : "Record Deleted Successfully",
                })
            }
        } catch (error) {
            console.log(error);
        }
    }

    return (
        <>
            <Container>
                <Row>
                    <h1>Are you sure you want to delete this record?</h1>
                    <Col xs="12" className="py-5">
                        <button className="btn btn-danger py-2" onClick={handleDelete}>Proceed</button>    
                    </Col>
                </Row>
            </Container>
        </>
    );
}

export default Delete