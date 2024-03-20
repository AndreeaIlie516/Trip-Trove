import { useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";
import ResponsiveAppBar from "../components/ResponsiveAppBar";
import Footer from "../components/Footer";
import {
  Box,
  Button,
  Typography,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
  Divider,
} from "@mui/material";
import { IDestination } from "../interfaces/Destination";
import { ILocation } from "../interfaces/Location";
import { useDestinations } from "../contexts/DestinationContext";
import { useLocations } from "../contexts/LocationContext";

export function Destination() {
  const { id } = useParams<{ id: string }>();
  const { getDestinationById, deleteDestination } = useDestinations();
  const { getLocationById } = useLocations();
  const [destination, setDestination] = useState<IDestination | undefined>();
  const [location, setLocation] = useState<ILocation | undefined>();
  const [open, setOpen] = useState(false);
  const [deleteId, setDeleteId] = useState<string | null>(null);

  const handleClickOpen = (id: string) => {
    console.log("Delete button clicked");
    setOpen(true);
    setDeleteId(id);
  };

  const handleClose = () => {
    console.log("Delete dialog closed");
    setOpen(false);
  };

  const handleDelete = () => {
    console.log("Delete dialog confirmed");
    if (deleteId) {
      deleteDestination(deleteId);
      setDeleteId(null);
    }
    setOpen(false);
  };

  useEffect(() => {
    const fetchDestinationDetails = async () => {
      if (!id) {
        return null;
      }
      const fetchedDestination = await getDestinationById(parseInt(id));
      setDestination(fetchedDestination);
    };
    fetchDestinationDetails();
    console.log(destination);

    const fetchLocationDetails = async () => {
      if (!destination) {
        return null;
      }
      if (!destination.location_id) {
        return null;
      }
      const fetchedLocation = await getLocationById(destination.location_id);
      setLocation(fetchedLocation);
    };
    fetchLocationDetails();
    console.log(location);
  });

  if (!destination || !location) {
    return (
      <Typography variant="h4" align="center">
        Loading...
      </Typography>
    );
  } else {
    return (
      <Box
        sx={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          justifyContent: "space-between",
          backgroundColor: "#f9f9f9",
          minHeight: "100vh",
          width: "100vw",
        }}
      >
        <ResponsiveAppBar />
        <Box
          sx={{
            borderRadius: "1em",
            overflow: "hidden",
            boxShadow: "0 10px 30px rgba(0,0,0,0.1)",
            width: "100%",
            maxWidth: "1080px",
          }}
        >
          <Box sx={{ width: "100%", overflow: "hidden" }}>
            <img
              src={destination.image_url}
              alt="Destination"
              style={{
                width: "100%",
                maxHeight: "500px",
                objectFit: "cover",
              }}
            />
          </Box>
          <Box
            sx={{
              p: 3,
              overflowY: "auto",
              maxHeight: "calc(100vh - 300px - 64px)",
            }}
          >
            <Typography
              variant="h4"
              component="h2"
              sx={{ textAlign: "center", mb: 2 }}
            >
              {destination.name}
            </Typography>
            <Divider sx={{ mb: 2 }} />
            <Typography variant="h6" component="h3" sx={{ mb: 1 }}>
              Location: {location.name}
            </Typography>
            <Typography variant="subtitle1" component="h3" sx={{ mb: 1 }}>
              Country: {location.country}
            </Typography>
            <Typography variant="body1" sx={{ mb: 2 }}>
              {destination.description}
            </Typography>
            <Box sx={{ display: "flex", justifyContent: "flex-end", mt: 2 }}>
              <Button
                variant="contained"
                color="primary"
                component={Link}
                to={`/destinations/update/${destination.id}`}
                sx={{ mr: 1 }}
              >
                Update
              </Button>
              <Button
                variant="contained"
                color="secondary"
                onClick={() => handleClickOpen(destination.id)}
              >
                Delete
              </Button>
            </Box>
          </Box>
        </Box>
        <Footer />
        <Dialog
          open={open}
          onClose={handleClose}
          aria-labelledby="alert-dialog-title"
          aria-describedby="alert-dialog-description"
        >
          <DialogTitle id="alert-dialog-title">
            {"Confirm Deletion"}
          </DialogTitle>
          <DialogContent>
            <DialogContentText id="alert-dialog-description">
              Are you sure you want to delete this destination?
            </DialogContentText>
          </DialogContent>
          <DialogActions>
            <Button onClick={handleClose}>Cancel</Button>
            <Button onClick={handleDelete} color="error">
              Delete
            </Button>
          </DialogActions>
        </Dialog>
      </Box>
    );
  }
}
