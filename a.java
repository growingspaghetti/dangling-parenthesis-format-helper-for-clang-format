    static Optional<String> getPreviousDC(int idInt) {
        if (idInt == 1) {
            return Optional.empty();
        }
        int    prev = idInt - 1;
        String d    = String.format("%04d_d", prev);
        if (new File("splits/" + d + ".jpg").exists()) {
            return Optional.of(d);
        }
        return Optional.of(String.format("%04d_c", prev));
    }
    
    static Optional<String> getPreviousDC(
        int idInt) {
        if (idInt == 1) {
            return Optional.empty();
        }
        int    prev = idInt - 1;
        String d    = String.format("%04d_d", prev);
        if (new File("splits/" + d + ".jpg").exists()) {
            return Optional.of(d);
        }
        return Optional.of(String.format("%04d_c", prev));
    }

    static Optional<String> getDaPreviousDC(
        int idInt) throws IllegaldadadaException {
        if (idInt == 1) {
            return Optional.empty();
        }
        int    prev = idInt - 1;
        String d    = String.format("%04d_d", prev);
        if (new File("splits/" + d + ".jpg").exists()) {
            return Optional.of(d);
        }
        return Optional.of(String.format("%04d_c", prev));
    }

    private static void fetchScreenShot(
        String url,
        String filePath) throws MalformedURLException,
                                IllegalStateException {
        if (new File(filePath).exists()) {
            return;
        }
        WebDriver driver = new RemoteWebDriver(
            new URL("http://127.0.0.1:4444/wd/hub"),
            DesiredCapabilities.chrome());
        driver.manage().window().setSize(new Dimension(1042, 1200)); //1042, 871));
        try {
            System.out.println(String.join(" -> ", url, filePath));
            driver.get(url);
            File source = ((TakesScreenshot)driver).getScreenshotAs(OutputType.FILE);
            System.out.println(source.getAbsolutePath());
            FileUtils.moveFile(source, new File(filePath));
            //FileUtils.copyFile(source, new File(filePath));
        } catch (Exception ex) {
            ex.printStackTrace();
        } finally {
            Optional.ofNullable(driver).ifPresent(WebDriver::quit); // not close
        }
    }