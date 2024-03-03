import React, { useState } from "react";
import Grid from "@mui/material/Grid";
import Paper from "@mui/material/Paper";
import DestinationCard from "./DestinationCard";
import Pagination from "@mui/material/Pagination";
import Box from "@mui/material/Box";
import { useNavigate } from "react-router-dom";
import { destinations } from "../mock/Data";

const DestinationGrid: React.FC = () => {
  const [page, setPage] = useState(1);
  const itemsPerPage = 9;
  const nav = useNavigate();

  return (
    <div>
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
        />
      </Box>
    </div>
  );
};

export default DestinationGrid;
