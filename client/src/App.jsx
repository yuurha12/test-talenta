import React, { useState, useEffect } from "react";
import { API } from "./config/api";
import {
  Form,
  Button,
  Row,
  Col,
  Table,
  Container,
  Navbar,
  Nav,
} from "react-bootstrap";
import Chart from "chart.js/auto";
import "./style/style.css";
import downloadPdf from "./download";
import { useQuery, useMutation } from "react-query";

const App = () => {
  const { data: friends, refetch: refetchFriends } = useQuery(
    "friendsCache",
    async () => {
      const response = await API.get("/friends");
      return response.data.data;
    }
  );

  const {
    data: statsData,
    isLoading: isLoadingStats,
    refetch: refetchStats,
  } = useQuery("friendStatsCache", async () => {
    const response = await API.get("/friendstats");
    return response.data.data;
  });

  const addFriendMutation = useMutation(async (friendData) => {
    const response = await API.post("/friend", friendData);
    return response.data.data;
  });

  const deleteFriendMutation = useMutation(async (id) => {
    await API.delete(`/friend/${id}`);
    return id;
  });

  const [chart, setChart] = useState(null);

  const handleFormSubmit = async (event) => {
    event.preventDefault();
    const formData = new FormData(event.target);
    const friendData = {
      name: formData.get("name"),
      gender: formData.get("gender"),
      age: parseInt(formData.get("age")),
    };
    try {
      await addFriendMutation.mutateAsync(friendData);
      refetchFriends();
      refetchStats();
    } catch (error) {
      console.log(error);
    }
  };

  const handleDeleteFriend = async (id) => {
    try {
      await deleteFriendMutation.mutateAsync(id);
      refetchFriends();
      refetchStats();
    } catch (error) {
      console.log(error);
    }
  };

  useEffect(() => {
    if (!statsData) return;

    const ctx = document.getElementById("chart");

    if (chart) {
      chart.destroy();
    }

    const newChart = new Chart(ctx, {
      type: "bar",
      data: {
        labels: ["Male", "Female"],
        datasets: [
          {
            label: "Gender",
            data: [statsData?.male_count, statsData.female_count],
            backgroundColor: ["#36A2EB", "#FF6384"],
            hoverBackgroundColor: ["#36A2EB", "#FF6384"],
          },
        ],
      },
      options: {
        scales: {
          y: {
            beginAtZero: true,
            stepSize: 1,
          },
        },
      },
    });

    setChart(newChart);

    return () => newChart.destroy();
  }, [statsData]);

  return (
    <>
      <Navbar bg="primary" variant="dark" className="nav">
        <Navbar.Brand href="#home">Friends App</Navbar.Brand>
        <Nav className="me-auto">
          <Nav.Link href="#stats">Friend Stats</Nav.Link>
          <Nav.Link href="#list">Friend List</Nav.Link>
          <Nav.Link href="#total">Friend Total</Nav.Link>
        </Nav>
      </Navbar>
      <Container>
        <Container>
            <div style={{marginTop: "50px"}}>

          <h1 id="home" className="title">
            Friends App
          </h1>
            </div>
          <Form onSubmit={handleFormSubmit} className="form">
            <Row>
              <Col>
                <Form.Group controlId="name">
                  <Form.Label>Name</Form.Label>
                  <Form.Control
                    type="text"
                    name="name"
                    placeholder="Enter your friend's name"
                    required
                  />
                </Form.Group>
              </Col>
              <Col>
                <Form.Group controlId="gender">
                  <Form.Label>Gender</Form.Label>
                  <Form.Control as="select" name="gender" required>
                    <option value="" disabled defaultValue>
                      Select Gender
                    </option>
                    <option value="male">Male</option>
                    <option value="female">Female</option>
                  </Form.Control>
                </Form.Group>
              </Col>
              <Col>
                <Form.Group controlId="age">
                  <Form.Label>Age</Form.Label>
                  <Form.Control
                    type="number"
                    name="age"
                    placeholder="Enter your friend's age"
                    required
                  />
                </Form.Group>
              </Col>
            </Row>
            <Button variant="primary" type="submit">
              Add Friend
            </Button>
          </Form>
          {!isLoadingStats && statsData && (
            <>
              <h2 id="stats" className="subtitle">Friend Stats</h2>
              <canvas id="chart" width="400px" height="400px"></canvas>
              <Button variant="success" onClick={downloadPdf}>
                Download PDF
              </Button>
            </>
          )}
          <Container id="list" className="fl">
            <h2 className="subtitle">Friends List</h2>
            <Table striped bordered hover className="my-table">
              <thead>
                <tr>
                  <th>#</th>
                  <th>Nama</th>
                  <th>Jenis Kelamin</th>
                  <th>Usia</th>
                  <th>Action</th>
                </tr>
              </thead>
              <tbody>
                {friends?.map((friend, index) => (
                  <tr key={friend.id}>
                    <td>{index + 1}</td>
                    <td>{friend.name}</td>
                    <td>{friend.gender}</td>
                    <td>{friend.age}</td>
                    <td>
                      <Button
                        variant="primary"
                        onClick={() => handleDeleteFriend(friend.id)}
                      >
                        Delete
                      </Button>
                    </td>
                  </tr>
                ))}
              </tbody>
            </Table>
            <div>
              <h2 id="total" className="details">Friends Total Details</h2>
            </div>
            {!isLoadingStats && statsData && (
              <Table striped bordered hover className="my-table">
                <thead>
                  <tr>
                    <th>Total Friend</th>
                    <th>Male Count</th>
                    <th>Female Count</th>
                    <th>Under 19 Count</th>
                    <th>Above 20 Count</th>
                  </tr>
                </thead>
                <tbody>
                  <tr>
                    <td>{statsData?.total_friend_count}</td>
                    <td>{statsData?.male_count}</td>
                    <td>{statsData?.female_count}</td>
                    <td>{statsData?.under_19_count}</td>
                    <td>{statsData?.above_20_count}</td>
                  </tr>
                </tbody>
              </Table>
            )}
          </Container>
        </Container>
      </Container>
    </>
  );
};

export default App;
