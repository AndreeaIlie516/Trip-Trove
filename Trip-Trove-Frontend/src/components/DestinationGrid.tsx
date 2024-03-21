import React, { useState, useMemo } from "react";
import DestinationCard from "./DestinationCard";
import {
  Box,
  Button,
  Grid,
  Pagination,
  Paper,
  Select,
  MenuItem,
  FormControl,
  InputLabel,
  SelectChangeEvent,
} from "@mui/material";
import { Link } from "react-router-dom";
import { useDestinations } from "../contexts/DestinationContext";
import { IDestination } from "../interfaces/Destination";

function sortDestinations(
  destinations: IDestination[],
  sortOrder: "recommended" | "asc" | "desc"
): IDestination[] {
  switch (sortOrder) {
    case "asc":
      return [...destinations].sort(
        (a, b) => a.visitors_last_year - b.visitors_last_year
      );
    case "desc":
      return [...destinations].sort(
        (a, b) => b.visitors_last_year - a.visitors_last_year
      );
    case "recommended":
    default:
      return destinations;
  }
}

const DestinationGrid: React.FC = () => {
  const { destinations } = useDestinations();
  const [page, setPage] = useState(1);
  const [itemsPerPage, setItemsPerPage] = useState(9);
  const [sortOrder, setSortOrder] = useState<"recommended" | "asc" | "desc">(
    "recommended"
  );

  const handleChange = (event: React.ChangeEvent<unknown>, value: number) => {
    setPage(value);
  };

  const handleItemsPerPageChange = (event: SelectChangeEvent<string>) => {
    const value = event.target.value;
    setItemsPerPage(parseInt(value, 10));
    setPage(1);
  };

  const handleChangePage = (
    event: React.ChangeEvent<unknown>,
    value: number
  ) => {
    setPage(value);
  };

  const sortedDestinations = useMemo(
    () => sortDestinations(destinations, sortOrder),
    [destinations, sortOrder]
  );

  const currentDestinations = useMemo(() => {
    const start = (page - 1) * itemsPerPage;
    return sortedDestinations.slice(start, start + itemsPerPage);
  }, [page, sortedDestinations, itemsPerPage]);

  const handleSortChange = (
    event: SelectChangeEvent<"recommended" | "asc" | "desc">
  ) => {
    setSortOrder(event.target.value as "recommended" | "asc" | "desc");
    setPage(1);
  };

  let accessLevel = 0;

  return (
    <>
      {accessLevel === 1 && (
        <Button
          component={Link}
          to={`/destinations/add`}
          variant="contained"
          style={{ marginTop: "20px", marginBottom: "40px" }}
        >
          Add New Private Destination
        </Button>
      )}
      <Box
        sx={{
          display: "flex",
          justifyContent: "center",
          marginTop: "2.2em",
          marginBottom: "2.2em",
        }}
      >
        <FormControl fullWidth>
          <InputLabel id="sort-order-select-label">Sort Order</InputLabel>
          <Select
            labelId="sort-order-select-label"
            id="sort-order-select"
            value={sortOrder}
            label="Sort Order"
            onChange={handleSortChange}
          >
            <MenuItem value="recommended">Recommended</MenuItem>
            <MenuItem value="asc">Ascending</MenuItem>
            <MenuItem value="desc">Descending</MenuItem>
          </Select>
        </FormControl>
      </Box>
      <Grid container columnSpacing={6} rowSpacing={6}>
        {currentDestinations
          .map((destination) => (
            <Grid item md={4} key={destination.id}>
              <Paper elevation={0}>
                <DestinationCard data={destination}></DestinationCard>
              </Paper>
            </Grid>
          ))}
      </Grid>
      <Box
        sx={{
          display: "flex",
          justifyContent: "center",
          marginTop: "2.2em",
        }}
      >
        <Pagination
          count={Math.ceil(sortedDestinations.length / itemsPerPage)}
          page={page}
          onChange={handleChange}
        />
        <FormControl style={{ marginTop: "-10px", width: "70px" }}>
          <Select
            labelId="items-per-page-label"
            id="items-per-page-select"
            value={itemsPerPage.toString()}
            label="Items Per Page"
            onChange={handleItemsPerPageChange}
          >
            <MenuItem value={9}>9</MenuItem>
            <MenuItem value={12}>12</MenuItem>
            <MenuItem value={15}>15</MenuItem>
            <MenuItem value={18}>18</MenuItem>
          </Select>
        </FormControl>
      </Box>
    </>
  );
};

export default DestinationGrid;
