import React, { useState } from "react";
import Grid from "@mui/material/Grid";
import Paper from "@mui/material/Paper";
import DestinationCard from "./DestinationCard";
import Pagination from "@mui/material/Pagination";
import { Box, Button } from "@mui/material";
import { Link, useNavigate } from "react-router-dom";
import { useDestinations } from "../contexts/DestinationContext";

const DestinationGrid: React.FC = () => {
  const { destinations } = useDestinations();
  const [page, setPage] = useState(1);
  const itemsPerPage = 9;
  const nav = useNavigate();

  const handleChange = (event: React.ChangeEvent<unknown>, value: number) => {
    setPage(value);
  };

  return (
    <div>
      <Button
        component={Link}
        to={`/destinations/add`}
        style={{ marginTop: "20px", marginBottom: "40px" }}
      >
        Add New Destination
      </Button>
      <Grid container columnSpacing={6} rowSpacing={6}>
        {destinations
          .slice((page - 1) * itemsPerPage, page * itemsPerPage)
          .map((destination) => (
            <Grid item md={4} key={destination.id}>
              <Paper elevation={0}>
                <DestinationCard data={destination}></DestinationCard>
              </Paper>
            </Grid>
          ))}
      </Grid>
      <Box
        sx={{ display: "flex", justifyContent: "center", marginTop: "2.2em" }}
      >
        <Pagination
          count={Math.ceil(destinations.length / itemsPerPage)}
          page={page}
          onChange={handleChange}
        />
      </Box>
    </div>
  );
};

export default DestinationGrid;
