import { useRef, useEffect } from "react";
import "./index.scss";

const App = () => {
  const mainRef = useRef<HTMLDivElement>(null);
  useEffect(() => {
    const setMinHeight = () => {
      if (!mainRef.current) return;
      mainRef.current.style.minHeight = `${window.innerHeight}px`;
    };
    setMinHeight();
    window.addEventListener("resize", setMinHeight);
    return () => {
      window.removeEventListener("resize", setMinHeight);
    };
  }, []);

  return (
    <div ref={mainRef} className="container">  
      <h1>GeoMeeting</h1>
    </div>
  );
};
export default App;
