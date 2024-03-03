import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import CardMedia from "@mui/material/CardMedia";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import { useNavigate } from "react-router-dom";
import { DestinationCardProps } from "../interfaces/Destination";

const DestinationCard: React.FC<DestinationCardProps> = ({ data }) => {
  const navigate = useNavigate();

  const handleClick = () => {
    let path = `/destinations/${data.id}`;
    navigate(path);
  };

  return (
    <Card
      sx={{
        borderRadius: "1em",
        border: "0px",
        boxShadow: "0 10px 4px rgba(0,0,0,0.2)",
        height: "420px",
        transition: "0.3s",
        ":hover": {
          boxShadow: "0 8px 16px rgba(0,0,0,0.25)",
          transform: "scale(1.05)",
          cursor: "pointer",
        },
      }}
    >
      <CardMedia
        component="img"
        alt="Destination"
        height="250"
        image={data.image_url}
      />
      <CardContent>
        <Typography
          gutterBottom
          variant="h5"
          component="div"
          sx={{ textAlign: "center" }}
        >
          {data.name}
          <br></br>
          <br></br>
          <Button variant="contained" onClick={handleClick}>
            See Details
          </Button>
        </Typography>
      </CardContent>
    </Card>
  );
};

export default DestinationCard;
