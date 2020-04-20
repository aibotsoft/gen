/* SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO"; */
/* SET AUTOCOMMIT = 0; */
/* START TRANSACTION; */
/* SET time_zone = "+00:00"; */

-- --------------------------------------------------------

--
-- Table structure for table `BettingStatusResponse` generated from model 'BettingStatusResponse'
--

CREATE TABLE IF NOT EXISTS `BettingStatusResponse` (
  `status` ENUM('ALL_BETTING_ENABLED', 'ALL_LIVE_BETTING_CLOSED', 'ALL_BETTING_CLOSED') NOT NULL COMMENT 'Betting status. '
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `CancellationDetailsItem` generated from model 'CancellationDetailsItem'
--

CREATE TABLE IF NOT EXISTS `CancellationDetailsItem` (
  `key` TEXT DEFAULT NULL,
  `value` TEXT DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `CancellationReason` generated from model 'CancellationReason'
-- Possible keys \\:   * correctTeam1Id * correctTeam2Id * correctListedPitcher1 * correctListedPitcher2 * correctSpread * correctTotalPoints * correctTeam1TotalPoints * correctTeam2TotalPoints * correctTeam1Score * correctTeam2Score * correctTeam1TennisSetsScore * correctTeam2TennisSetsScore 
--

CREATE TABLE IF NOT EXISTS `CancellationReason` (
  `code` TEXT NOT NULL,
  `details` TEXT DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Possible keys \\:   * correctTeam1Id * correctTeam2Id * correctListedPitcher1 * correctListedPitcher2 * correctSpread * correctTotalPoints * correctTeam1TotalPoints * correctTeam2TotalPoints * correctTeam1Score * correctTeam2Score * correctTeam1TennisSetsScore * correctTeam2TennisSetsScore ';

--
-- Table structure for table `CancellationReasonDetailsType` generated from model 'CancellationReasonDetailsType'
--

CREATE TABLE IF NOT EXISTS `CancellationReasonDetailsType` (
  `key` TEXT DEFAULT NULL,
  `value` TEXT DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `CancellationReasonResponse` generated from model 'CancellationReasonResponse'
-- Cancellation Response Data
--

CREATE TABLE IF NOT EXISTS `CancellationReasonResponse` (
  `cancellationReasons` TEXT DEFAULT NULL COMMENT 'Contains a list of Cancellation Reasons.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Cancellation Response Data';

--
-- Table structure for table `CancellationReasonType` generated from model 'CancellationReasonType'
--

CREATE TABLE IF NOT EXISTS `CancellationReasonType` (
  `code` TEXT DEFAULT NULL COMMENT 'Cancellation Reason Code',
  `details` TEXT DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `ClientBalanceResponse` generated from model 'ClientBalanceResponse'
-- Client Balance Details
--

CREATE TABLE IF NOT EXISTS `ClientBalanceResponse` (
  `availableBalance` DECIMAL(20, 9) NOT NULL COMMENT 'Amount available for betting.',
  `outstandingTransactions` DECIMAL(20, 9) NOT NULL COMMENT 'Sum of not yet settled bet amounts.',
  `givenCredit` DECIMAL(20, 9) NOT NULL COMMENT 'Client’s credit.',
  `currency` TEXT NOT NULL COMMENT 'Client’s currency code.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Client Balance Details';

--
-- Table structure for table `Currency` generated from model 'Currency'
--

CREATE TABLE IF NOT EXISTS `Currency` (
  `code` TEXT DEFAULT NULL COMMENT 'Currency code.',
  `name` TEXT DEFAULT NULL COMMENT 'Currency name.',
  `rate` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Exchange rate to USD.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `ErrorResponse` generated from model 'ErrorResponse'
--

CREATE TABLE IF NOT EXISTS `ErrorResponse` (
  `code` ENUM('INVALID_REQUEST_DATA', 'INVALID_CREDENTIALS', 'INVALID_AUTHORIZATION_HEADER', 'ACCOUNT_INACTIVE', 'NO_API_ACCESS', 'RESUBMIT_REQUEST') DEFAULT NULL COMMENT 'INVALID_REQUEST_DATA  &#x3D; Invalid request parameters (http status 400)   INVALID_CREDENTIALS &#x3D; Authorization failed, invalid credentials (http status 401)   INVALID_AUTHORIZATION_HEADER &#x3D; HTTP Authorization header is missing (http status 401)  ACCOUNT_INACTIVE &#x3D; Client&#39;s account is not active (http status 403)   NO_API_ACCESS &#x3D; Account not permitted to access the API (http status 403)  RESUBMIT_REQUEST &#x3D; It can happen only when placing a bet (http status 400).  Unable to process the request but the request itself is valid. This happens more often on the live betting in situations when there is more than one place bet request at the same on the same line. When this happens, we don&#39;t keep the place bet request on the server until we know if we can accept or reject the bet, but instead we return the error. It&#39;s also very likely that the line will change after that. To reduce a chance of getting RESUBMIT_REQUEST client should try to place a bet as fast as possible.  ',
  `message` TEXT DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `ErrorResponseWithErrorRef` generated from model 'ErrorResponseWithErrorRef'
--

CREATE TABLE IF NOT EXISTS `ErrorResponseWithErrorRef` (
  `ref` TEXT DEFAULT NULL,
  `code` TEXT DEFAULT NULL,
  `message` TEXT DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `ExtendedErrorResponse` generated from model 'ExtendedErrorResponse'
--

CREATE TABLE IF NOT EXISTS `ExtendedErrorResponse` (
  `ref` TEXT DEFAULT NULL,
  `code` TEXT DEFAULT NULL,
  `message` TEXT DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `ExtendedLinesErrorResponse` generated from model 'ExtendedLinesErrorResponse'
--

CREATE TABLE IF NOT EXISTS `ExtendedLinesErrorResponse` (
  `ref` TEXT DEFAULT NULL,
  `status` TEXT DEFAULT NULL,
  `error` TEXT DEFAULT NULL,
  `code` INT DEFAULT NULL COMMENT 'Code identifying an error that occurred.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `Fixture` generated from model 'Fixture'
--

CREATE TABLE IF NOT EXISTS `Fixture` (
  `id` BIGINT DEFAULT NULL COMMENT 'Event id.',
  `parentId` BIGINT DEFAULT NULL COMMENT 'If event is linked to another event, parentId will be populated.  Live event would have pre game event as parent id.',
  `starts` DATETIME DEFAULT NULL COMMENT 'Start time of the event in UTC.',
  `home` TEXT DEFAULT NULL COMMENT 'Home team name.',
  `away` TEXT DEFAULT NULL COMMENT 'Away team name.',
  `rotNum` TEXT DEFAULT NULL COMMENT 'Team1 rotation number. Please note that in the next version of /fixtures, rotNum property will be decommissioned. ParentId can be used instead to group the related events.',
  `liveStatus` ENUM('0', '1', '2') DEFAULT NULL COMMENT 'Indicates live status of the event.   0 &#x3D; No live betting will be offered on this event,  1 &#x3D; Live betting event,  2 &#x3D; Live betting will be offered on this match, but on a different event. Please note that [pre-game and live events are different](https://github.com/pinnacleapi/pinnacleapi-documentation/blob/master/FAQ.md#how-to-find-associated-events) . ',
  `homePitcher` TEXT DEFAULT NULL COMMENT 'Home team pitcher. Only for Baseball.',
  `awayPitcher` TEXT DEFAULT NULL COMMENT 'Away team pitcher. Only for Baseball.',
  `status` ENUM('O', 'H', 'I') DEFAULT NULL COMMENT 'This is deprecated parameter, please check period&#39;s &#x60;status&#x60; in the &#x60;/odds&#x60; endpoint to see if it&#39;s open for betting.   O &#x3D; This is the starting status of a game.    H &#x3D; This status indicates that the lines are temporarily unavailable for betting,   I &#x3D; This status indicates that one or more lines have a red circle (lower maximum bet amount). ',
  `parlayRestriction` ENUM('0', '1', '2') DEFAULT NULL COMMENT ' Parlay status of the event.   0 &#x3D; Allowed to parlay, without restrictions,  1 &#x3D; Not allowed to parlay this event,  2 &#x3D; Allowed to parlay with the restrictions. You cannot have more than one leg from the same event in the parlay. All events with the same rotation number are treated as same event. ',
  `altTeaser` TINYINT(1) DEFAULT NULL COMMENT 'Whether an event is offer with alternative teaser points. Events with alternative teaser points may vary from teaser definition.',
  `resultingUnit` TEXT DEFAULT NULL COMMENT 'Specifies based on what the event will be resulted, e.g. Corners, Bookings  '
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `FixturesLeague` generated from model 'FixturesLeague'
--

CREATE TABLE IF NOT EXISTS `FixturesLeague` (
  `id` INT DEFAULT NULL COMMENT 'League ID.',
  `name` TEXT DEFAULT NULL COMMENT 'League Name.',
  `events` TEXT DEFAULT NULL COMMENT 'Contains a list of events.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `FixturesResponse` generated from model 'FixturesResponse'
--

CREATE TABLE IF NOT EXISTS `FixturesResponse` (
  `sportId` INT DEFAULT NULL COMMENT 'Same as requested sport Id.',
  `last` BIGINT DEFAULT NULL COMMENT 'Use this value for the subsequent requests for since query parameter to get just the changes since previous response.',
  `league` TEXT DEFAULT NULL COMMENT 'Contains a list of Leagues.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `GetBetsByTypeResponseV3` generated from model 'GetBetsByTypeResponseV3'
--

CREATE TABLE IF NOT EXISTS `GetBetsByTypeResponseV3` (
  `moreAvailable` TINYINT(1) DEFAULT NULL COMMENT 'Whether there are more pages available.',
  `pageSize` INT DEFAULT NULL COMMENT 'Page size. Default is 1000.',
  `fromRecord` INT DEFAULT NULL COMMENT 'Starting record number of the result set. Records start at zero',
  `toRecord` INT DEFAULT NULL COMMENT 'Ending record number of the result set.',
  `straightBets` TEXT DEFAULT NULL COMMENT 'A collection of placed straight bets.',
  `parlayBets` TEXT DEFAULT NULL COMMENT 'A collection of placed parlay bets.',
  `teaserBets` TEXT DEFAULT NULL COMMENT 'A collection of placed teaser bets.',
  `specialBets` TEXT DEFAULT NULL COMMENT 'A collection of placed special bets.',
  `manualBets` TEXT DEFAULT NULL COMMENT 'A collection of placed manual bets.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `InRunningEvent` generated from model 'InRunningEvent'
--

CREATE TABLE IF NOT EXISTS `InRunningEvent` (
  `id` BIGINT DEFAULT NULL COMMENT 'Game Id',
  `state` ENUM('1', '2', '3', '4', '5', '6', '7', '8', '9', '10', '11') DEFAULT NULL COMMENT 'State of the game.  1 &#x3D; First half in progress,  2 &#x3D; Half time in progress,  3 &#x3D; Second half in progress,  4 &#x3D; End of regular time, 5 &#x3D; First half extra time in progress,  6 &#x3D; Extra time half time in progress,  7 &#x3D; Second half extra time in progress,  8 &#x3D; End of extra time,  9 &#x3D; End of Game,  10 &#x3D; Game is temporary suspended,  11 &#x3D; Penalties in progress ',
  `elapsed` INT DEFAULT NULL COMMENT 'Elapsed minutes'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `InRunningLeague` generated from model 'InRunningLeague'
--

CREATE TABLE IF NOT EXISTS `InRunningLeague` (
  `id` INT DEFAULT NULL COMMENT 'League Id',
  `events` TEXT DEFAULT NULL COMMENT 'Events container'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `InRunningResponse` generated from model 'InRunningResponse'
--

CREATE TABLE IF NOT EXISTS `InRunningResponse` (
  `sports` TEXT DEFAULT NULL COMMENT 'Sports container'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `InRunningSport` generated from model 'InRunningSport'
--

CREATE TABLE IF NOT EXISTS `InRunningSport` (
  `id` INT DEFAULT NULL COMMENT 'Sport Id',
  `leagues` TEXT DEFAULT NULL COMMENT 'Leagues container'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `League` generated from model 'League'
--

CREATE TABLE IF NOT EXISTS `League` (
  `id` INT DEFAULT NULL COMMENT 'League Id.',
  `name` TEXT DEFAULT NULL COMMENT 'Name of the league.',
  `homeTeamType` TEXT DEFAULT NULL COMMENT 'Specifies whether the home team is team1 or team2. You need this information to place a bet.',
  `hasOfferings` TINYINT(1) DEFAULT NULL COMMENT 'Whether the league currently has events or specials.',
  `container` TEXT DEFAULT NULL COMMENT 'Represents grouping for the league, usually a region/country',
  `allowRoundRobins` TINYINT(1) DEFAULT NULL COMMENT 'Specifies whether you can place parlay round robins on events in this league.',
  `leagueSpecialsCount` INT DEFAULT NULL COMMENT 'Indicates how many specials are in the given league.',
  `eventSpecialsCount` INT DEFAULT NULL COMMENT 'Indicates how many game specials are in the given league.',
  `eventCount` INT DEFAULT NULL COMMENT 'Indicates how many events are in the given league.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `Leagues` generated from model 'Leagues'
--

CREATE TABLE IF NOT EXISTS `Leagues` (
  `leagues` TEXT DEFAULT NULL COMMENT 'Leagues container'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `LineResponse` generated from model 'LineResponse'
--

CREATE TABLE IF NOT EXISTS `LineResponse` (
  `status` ENUM('SUCCESS', 'NOT_EXISTS') DEFAULT NULL COMMENT 'If the value is NOT_EXISTS, than this will be the only parameter in the response. All other params would be empty. [SUCCESS &#x3D; OK, NOT_EXISTS &#x3D; Line not offered anymore]',
  `price` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Latest price.',
  `lineId` BIGINT DEFAULT NULL COMMENT 'Line identification needed to place a bet.',
  `altLineId` BIGINT DEFAULT NULL COMMENT 'This would be needed to place the bet if the handicap is on alternate line, otherwise it will not be populated in the response.',
  `team1Score` INT DEFAULT NULL COMMENT 'Away team score. Applicable to soccer only.',
  `team2Score` INT DEFAULT NULL COMMENT 'Home team score. Applicable to soccer only.',
  `team1RedCards` INT DEFAULT NULL COMMENT 'Team 1 red cards. Applicable to soccer only.',
  `team2RedCards` INT DEFAULT NULL COMMENT 'Team 2 red cards. Applicable to soccer only.',
  `maxRiskStake` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Maximum bettable risk amount.',
  `minRiskStake` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Minimum bettable risk amount.',
  `maxWinStake` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Maximum bettable win amount.',
  `minWinStake` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Minimum bettable win amount.',
  `effectiveAsOf` TEXT DEFAULT NULL COMMENT 'Line is effective as of this date and time in UTC.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `LinesErrorResponse` generated from model 'LinesErrorResponse'
--

CREATE TABLE IF NOT EXISTS `LinesErrorResponse` (
  `status` TEXT DEFAULT NULL,
  `error` TEXT DEFAULT NULL,
  `code` INT DEFAULT NULL COMMENT 'Code identifying an error that occurred.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `ManualBet` generated from model 'ManualBet'
--

CREATE TABLE IF NOT EXISTS `ManualBet` (
  `betId` BIGINT NOT NULL COMMENT 'Bet identification',
  `wagerNumber` INT NOT NULL COMMENT 'Wager identification. All bets placed thru the API will have value 1. Website Classic view supports multiple contest(special) bets placement in the same bet slip in that case the bet would have appropriate wager number, as well as all round robin parlay bets.',
  `placedAt` DATETIME NOT NULL COMMENT 'Date time when the bet was placed.',
  `betStatus` ENUM('ACCEPTED', 'CANCELLED', 'LOSE', 'REFUNDED', 'WON') NOT NULL COMMENT 'Bet Status.   ACCEPTED &#x3D; Bet was accepted,   CANCELLED &#x3D; Bet is cancelled as per Pinnacle betting rules,   LOSE &#x3D; The bet is settled as lose,   REFUNDED &#x3D; When an event is cancelled or when the bet is settled as push, the bet will have REFUNDED status,   WON &#x3D; The bet is settled as won  ',
  `betType` TEXT NOT NULL,
  `win` DECIMAL(20, 9) NOT NULL COMMENT 'Win amount.',
  `risk` DECIMAL(20, 9) NOT NULL COMMENT 'Risk amount.',
  `winLoss` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Win-Loss for settled bets.',
  `updateSequence` BIGINT NOT NULL COMMENT 'Update Sequence',
  `description` TEXT NOT NULL COMMENT 'Manual bet description.',
  `referenceBetId` BIGINT DEFAULT NULL COMMENT 'Referenced original bet id.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `MultiBetRequest.SpecialBetRequest` generated from model 'MultiBetRequestPeriodSpecialBetRequest'
--

CREATE TABLE IF NOT EXISTS `MultiBetRequest.SpecialBetRequest` (
  `bets` TEXT DEFAULT NULL COMMENT 'The individual bets.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `MultiBetResponse.SpecialBetResponse` generated from model 'MultiBetResponsePeriodSpecialBetResponse'
--

CREATE TABLE IF NOT EXISTS `MultiBetResponse.SpecialBetResponse` (
  `bets` TEXT DEFAULT NULL COMMENT 'The individual bets.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `OddsEvent` generated from model 'OddsEvent'
--

CREATE TABLE IF NOT EXISTS `OddsEvent` (
  `id` BIGINT DEFAULT NULL COMMENT 'Event Id.',
  `awayScore` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Away team score. Only for live soccer events.Supported only for full match period (number&#x3D;0).',
  `homeScore` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Home team score. Only for live soccer events.Supported only for full match period (number&#x3D;0).',
  `awayRedCards` INT DEFAULT NULL COMMENT 'Away team red cards. Only for live soccer events. Supported only for full match period (number&#x3D;0).',
  `homeRedCards` INT DEFAULT NULL COMMENT 'Home team red cards. Only for live soccer events.Supported only for full match period (number&#x3D;0).',
  `periods` TEXT DEFAULT NULL COMMENT 'Contains a list of periods.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `OddsLeague` generated from model 'OddsLeague'
--

CREATE TABLE IF NOT EXISTS `OddsLeague` (
  `id` INT DEFAULT NULL COMMENT 'League Id.',
  `events` TEXT DEFAULT NULL COMMENT 'Contains a list of events.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `OddsMoneyline` generated from model 'OddsMoneyline'
--

CREATE TABLE IF NOT EXISTS `OddsMoneyline` (
  `home` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Away team price',
  `away` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Away team price.',
  `draw` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Draw price. This is present only for events we offer price for draw.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `OddsPeriod` generated from model 'OddsPeriod'
--

CREATE TABLE IF NOT EXISTS `OddsPeriod` (
  `lineId` BIGINT DEFAULT NULL COMMENT 'Line Id.',
  `number` INT DEFAULT NULL COMMENT 'This represents the period of the match. For example, for soccer we have  0 (Game), 1 (1st Half) &amp; 2 (2nd Half)',
  `cutoff` DATETIME DEFAULT NULL COMMENT 'Period’s wagering cut-off date in UTC.',
  `status` INT DEFAULT NULL COMMENT '1 - online, period is open for betting  2 - offline, period is not open for betting ',
  `maxSpread` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Maximum spread bet volume. See [How to calculate max risk from the max volume](https://github.com/pinnacleapi/pinnacleapi-documentation/blob/master/FAQ.md#how-to-calculate-max-risk-from-the-max-volume-limits-in-odds)',
  `maxMoneyline` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Maximum moneyline bet volume. See [How to calculate max risk from the max volume](https://github.com/pinnacleapi/pinnacleapi-documentation/blob/master/FAQ.md#how-to-calculate-max-risk-from-the-max-volume-limits-in-odds)',
  `maxTotal` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Maximum total points bet volume. See [How to calculate max risk from the max volume](https://github.com/pinnacleapi/pinnacleapi-documentation/blob/master/FAQ.md#how-to-calculate-max-risk-from-the-max-volume-limits-in-odds)',
  `maxTeamTotal` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Maximum team total points bet volume. See [How to calculate max risk from the max volume](https://github.com/pinnacleapi/pinnacleapi-documentation/blob/master/FAQ.md#how-to-calculate-max-risk-from-the-max-volume-limits-in-odds)',
  `spreads` TEXT DEFAULT NULL COMMENT 'Container for spread odds.',
  `moneyline` TEXT DEFAULT NULL,
  `totals` TEXT DEFAULT NULL COMMENT 'Container for team total points.',
  `teamTotal` TEXT DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `OddsResponse` generated from model 'OddsResponse'
--

CREATE TABLE IF NOT EXISTS `OddsResponse` (
  `sportId` INT DEFAULT NULL COMMENT 'Same as requested sport Id.',
  `last` BIGINT DEFAULT NULL COMMENT 'Use this value for the subsequent requests for since query parameter to get just the changes since previous response.',
  `leagues` TEXT DEFAULT NULL COMMENT 'Contains a list of Leagues.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `OddsSpread` generated from model 'OddsSpread'
--

CREATE TABLE IF NOT EXISTS `OddsSpread` (
  `altLineId` BIGINT DEFAULT NULL COMMENT 'This is present only if it’s alternative line.',
  `hdp` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Home team handicap.',
  `home` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Home team price.',
  `away` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Away team price.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `OddsTeamTotal` generated from model 'OddsTeamTotal'
--

CREATE TABLE IF NOT EXISTS `OddsTeamTotal` (
  `points` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Total points.',
  `over` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Over price.',
  `under` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Under price.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `OddsTeamTotals` generated from model 'OddsTeamTotals'
--

CREATE TABLE IF NOT EXISTS `OddsTeamTotals` (
  `home` TEXT DEFAULT NULL,
  `away` TEXT DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `OddsTotal` generated from model 'OddsTotal'
--

CREATE TABLE IF NOT EXISTS `OddsTotal` (
  `altLineId` BIGINT DEFAULT NULL COMMENT 'This is present only if it’s alternative line.',
  `points` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Total points.',
  `over` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Over price.',
  `under` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Under price.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `ParlayBet` generated from model 'ParlayBet'
--

CREATE TABLE IF NOT EXISTS `ParlayBet` (
  `betId` BIGINT NOT NULL COMMENT 'Bet identification',
  `uniqueRequestId` TEXT DEFAULT NULL COMMENT 'Unique Request Id',
  `wagerNumber` INT NOT NULL COMMENT 'Wager identification. All bets placed thru the API will have value 1. Website Classic view supports multiple contest(special) bets placement in the same bet slip in that case the bet would have appropriate wager number, as well as all round robin parlay bets.',
  `placedAt` DATETIME NOT NULL COMMENT 'Date time when the bet was placed.',
  `betStatus` ENUM('ACCEPTED', 'CANCELLED', 'LOSE', 'PENDING_ACCEPTANCE', 'REFUNDED', 'NOT_ACCEPTED', 'WON') NOT NULL COMMENT 'Bet Status.   ACCEPTED &#x3D; Bet was accepted,   CANCELLED &#x3D; Bet is cancelled as per Pinnacle betting rules,   LOSE &#x3D; The bet is settled as lose,   PENDING_ACCEPTANCE &#x3D; This status is reserved only for live bets. If a live bet is placed during danger zone or live delay is applied, it will be in PENDING_ACCEPTANCE , otherwise in ACCEPTED status. From this status bet can go to ACCEPTED or REJECTED status,   REFUNDED &#x3D; When an event is cancelled or when the bet is settled as push, the bet will have REFUNDED status,   NOT_ACCEPTED &#x3D; Bet was not accepted. Bet can be in this status only if it was previously in PENDING_ACCEPTANCE status,   WON &#x3D; The bet is settled as won ',
  `betType` TEXT NOT NULL,
  `win` DECIMAL(20, 9) NOT NULL COMMENT 'Win amount.',
  `risk` DECIMAL(20, 9) NOT NULL COMMENT 'Risk amount.',
  `winLoss` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Win-Loss for settled bets.',
  `oddsFormat` TEXT NOT NULL,
  `customerCommission` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Client’s commission on the bet.',
  `cancellationReason` TEXT DEFAULT NULL,
  `updateSequence` BIGINT NOT NULL COMMENT 'Update Sequence',
  `legs` TEXT NOT NULL,
  `price` DECIMAL(20, 9) DEFAULT NULL,
  `finalPrice` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Only for settled parlay. Final price may differ in case leg was cancelled or half won'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `ParlayLeg` generated from model 'ParlayLeg'
--

CREATE TABLE IF NOT EXISTS `ParlayLeg` (
  `sportId` INT DEFAULT NULL,
  `legBetType` ENUM('MONEYLINE', 'SPREAD', 'TOTAL_POINTS') DEFAULT NULL COMMENT 'Parlay leg type.',
  `legBetStatus` ENUM('CANCELLED', 'LOSE', 'PUSH', 'REFUNDED', 'WON', 'ACCEPTED') DEFAULT NULL COMMENT 'Parlay Leg status. CANCELLED &#x3D; The leg is canceled- the stake on this leg will be transferred to the next one. In this case the leg will be ignored when calculating the winLoss,   LOSE &#x3D; The leg is a loss or a push-lose. When Push-lose happens, the half of the stake on the leg will be pushed to the next leg, and the other half will be a lose. This can happen only when the leg is placed on a quarter points handicap,   PUSH &#x3D; The leg is a push - the stake on this leg will be transferred to the next one. In this case the leg will be ignored when calculating the winLoss,   REFUNDED &#x3D; The leg is refunded - the stake on this leg will be transferred to the next one. In this case the leg will be ignored when calculating the winLoss,   WON &#x3D; The leg is a won or a push-won. When Push-won happens, the half of the stake on the leg will be pushed to the next leg, and the other half is won. This can happen only when the leg is placed on a quarter points handicap  ',
  `leagueId` INT DEFAULT NULL,
  `eventId` BIGINT DEFAULT NULL,
  `eventStartTime` DATETIME DEFAULT NULL COMMENT 'Date time when the event starts.',
  `handicap` DECIMAL(20, 9) DEFAULT NULL,
  `price` DECIMAL(20, 9) DEFAULT NULL,
  `teamName` TEXT DEFAULT NULL,
  `side` ENUM('OVER', 'UNDER') DEFAULT NULL COMMENT 'Side type.',
  `pitcher1` TEXT DEFAULT NULL,
  `pitcher2` TEXT DEFAULT NULL,
  `pitcher1MustStart` TINYINT(1) DEFAULT NULL,
  `pitcher2MustStart` TINYINT(1) DEFAULT NULL,
  `team1` TEXT DEFAULT NULL COMMENT 'Wellington Phoenix',
  `team2` TEXT DEFAULT NULL COMMENT 'Adelaide United',
  `periodNumber` INT DEFAULT NULL,
  `ftTeam1Score` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Full time team 1 score',
  `ftTeam2Score` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Full time team 2 score',
  `pTeam1Score` DECIMAL(20, 9) DEFAULT NULL COMMENT 'End of period team 1 score. If the bet was placed on Game period (periodNumber &#x3D;0) , this will be null',
  `pTeam2Score` DECIMAL(20, 9) DEFAULT NULL COMMENT 'End of period team 2 score. If the bet was placed on Game period (periodNumber &#x3D;0) , this will be null',
  `cancellationReason` TEXT DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `PlaceBetRequest` generated from model 'PlaceBetRequest'
-- Request to place a bet.
--

CREATE TABLE IF NOT EXISTS `PlaceBetRequest` (
  `oddsFormat` TEXT DEFAULT NULL,
  `uniqueRequestId` TEXT DEFAULT NULL COMMENT 'This is a Unique ID for PlaceBet requests. This is to support idempotent requests.',
  `acceptBetterLine` TINYINT(1) DEFAULT NULL COMMENT 'Whether or not to accept a bet when there is a line change in favor of the client.',
  `stake` DECIMAL(20, 9) DEFAULT NULL COMMENT 'amount in client’s currency.',
  `winRiskStake` ENUM('WIN', 'RISK') DEFAULT NULL COMMENT 'Whether the stake amount is risk or win amount.',
  `lineId` BIGINT DEFAULT NULL COMMENT 'Line identification.',
  `altLineId` BIGINT DEFAULT NULL COMMENT 'Alternate line identification.',
  `pitcher1MustStart` TINYINT(1) DEFAULT NULL COMMENT 'Baseball only. Refers to the pitcher for Team1. This applicable only for MONEYLINE bet type, for all other bet types this has to be TRUE.',
  `pitcher2MustStart` TINYINT(1) DEFAULT NULL COMMENT 'Baseball only. Refers to the pitcher for Team2. This applicable only for MONEYLINE bet type, for all other bet types this has to be TRUE.',
  `fillType` ENUM('NORMAL', 'FILLANDKILL', 'FILLMAXLIMIT') DEFAULT 'NORMAL' COMMENT 'NORMAL - bet will be placed on specified stake.   FILLANDKILL - If the stake is over the max limit, bet will be placed on max limit, otherwise it will be placed on specified stake.   FILLMAXLIMIT - bet will be places on max limit, stake amount will be ignored. Please note that maximum limits can change at any moment, which may result in risking more than anticipated. This option is replacement of isMaxStakeBet from v1/bets/place&#39; ',
  `sportId` INT DEFAULT NULL,
  `eventId` BIGINT DEFAULT NULL,
  `periodNumber` INT DEFAULT NULL,
  `betType` ENUM('MONEYLINE', 'TEAM_TOTAL_POINTS', 'SPREAD', 'TOTAL_POINTS') DEFAULT NULL COMMENT 'Bet type.',
  `team` ENUM('TEAM1', 'TEAM2', 'DRAW') DEFAULT NULL COMMENT 'Team type.',
  `side` ENUM('OVER', 'UNDER') DEFAULT NULL COMMENT 'Side type.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Request to place a bet.';

--
-- Table structure for table `PlaceBetResponseV2` generated from model 'PlaceBetResponseV2'
--

CREATE TABLE IF NOT EXISTS `PlaceBetResponseV2` (
  `status` ENUM('ACCEPTED', 'PENDING_ACCEPTANCE', 'PROCESSED_WITH_ERROR') DEFAULT NULL COMMENT 'Status of the response.',
  `errorCode` ENUM('ALL_BETTING_CLOSED', 'ALL_LIVE_BETTING_CLOSED', 'ABOVE_EVENT_MAX', 'ABOVE_MAX_BET_AMOUNT', 'BELOW_MIN_BET_AMOUNT', 'BLOCKED_BETTING', 'BLOCKED_CLIENT', 'INSUFFICIENT_FUNDS', 'INVALID_COUNTRY', 'INVALID_EVENT', 'INVALID_ODDS_FORMAT', 'LINE_CHANGED', 'LISTED_PITCHERS_SELECTION_ERROR', 'OFFLINE_EVENT', 'PAST_CUTOFFTIME', 'RED_CARDS_CHANGED', 'SCORE_CHANGED', 'TIME_RESTRICTION', 'DUPLICATE_UNIQUE_REQUEST_ID', 'INCOMPLETE_CUSTOMER_BETTING_PROFILE', 'INVALID_CUSTOMER_PROFILE', 'LIMITS_CONFIGURATION_ISSUE', 'RESPONSIBLE_BETTING_LOSS_LIMIT_EXCEEDED', 'RESPONSIBLE_BETTING_RISK_LIMIT_EXCEEDED', 'SYSTEM_ERROR_3', 'LICENCE_RESTRICTION_LIVE_BETTING_BLOCKED') DEFAULT NULL COMMENT 'If Status is PROCESSED_WITH_ERROR, errorCode will be in the response.   ALL_BETTING_CLOSED &#x3D; Betting is not allowed at this moment. This may happen during system maintenance,   ALL_LIVE_BETTING_CLOSED &#x3D; Live betting is not allowed at this moment. This may happen during system maintenance,   ABOVE_EVENT_MAX &#x3D; Bet cannot be placed because client exceeded allowed maximum of risk on a line,   ABOVE_MAX_BET_AMOUNT &#x3D; Stake is above allowed maximum amount,    BELOW_MIN_BET_AMOUNT &#x3D; Stake is below allowed minimum amount,   BLOCKED_BETTING &#x3D; Betting is suspended for the client,   BLOCKED_CLIENT &#x3D; Client is no longer active,    INSUFFICIENT_FUNDS &#x3D; Bet is submitted by a client with insufficient funds,   INVALID_COUNTRY &#x3D; Client country is not allowed for betting,   INVALID_EVENT &#x3D; Invalid eventid,   INVALID_ODDS_FORMAT &#x3D; If a bet was submitted with the odds format that is not allowed for the client,   LINE_CHANGED &#x3D; Bet is submitted on a line that has changed,   LISTED_PITCHERS_SELECTION_ERROR &#x3D; If bet was submitted with pitcher1MustStart and/or pitcher2MustStart parameters in Place Bet request with values that are not allowed,   OFFLINE_EVENT &#x3D; Bet is submitted on an event that is offline or the submitted line is not offered at the moment due to points/handicap change or the submitted bet type is just not offered at the moment,   PAST_CUTOFFTIME &#x3D; Bet is submitted on a game after the betting cutoff time,   RED_CARDS_CHANGED &#x3D; Bet is submitted on a live soccer event with changed red card count,   SCORE_CHANGED &#x3D; Bet is submitted on a live soccer event with changed score,   TIME_RESTRICTION &#x3D; Bet is submitted within too short of a period from the same bet previously placed by a client,   DUPLICATE_UNIQUE_REQUEST_ID &#x3D; Request with the same uniqueRequestId was already processed. Please set the new value if you still want the request to be processed,   INCOMPLETE_CUSTOMER_BETTING_PROFILE &#x3D; System configuration issue,   INVALID_CUSTOMER_PROFILE &#x3D; System configuration issue,   LIMITS_CONFIGURATION_ISSUE &#x3D; System configuration issue,   RESPONSIBLE_BETTING_LOSS_LIMIT_EXCEEDED &#x3D; Client has reached his total loss limit,   RESPONSIBLE_BETTING_RISK_LIMIT_EXCEEDED &#x3D; Client has reached his total risk limit,   SYSTEM_ERROR_3 &#x3D; Unexpected error,   LICENCE_RESTRICTION_LIVE_BETTING_BLOCKED - Live betting blocked due to licence restrictions ',
  `uniqueRequestId` TEXT DEFAULT NULL COMMENT 'Echo of the uniqueRequestId from the request.',
  `straightBet` TEXT DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `SettledFixturesEvent` generated from model 'SettledFixturesEvent'
--

CREATE TABLE IF NOT EXISTS `SettledFixturesEvent` (
  `id` BIGINT DEFAULT NULL COMMENT 'Event Id.',
  `periods` TEXT DEFAULT NULL COMMENT 'Contains a list of periods.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `SettledFixturesLeague` generated from model 'SettledFixturesLeague'
--

CREATE TABLE IF NOT EXISTS `SettledFixturesLeague` (
  `id` INT DEFAULT NULL COMMENT 'League Id.',
  `events` TEXT DEFAULT NULL COMMENT 'Contains a list of events.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `SettledFixturesPeriod` generated from model 'SettledFixturesPeriod'
--

CREATE TABLE IF NOT EXISTS `SettledFixturesPeriod` (
  `number` INT DEFAULT NULL COMMENT 'This represents the period of the match. For example, for soccer we have 0 (Game), 1 (1st Half) &amp; 2 (2nd Half)',
  `status` ENUM('1', '2', '3', '4', '5') DEFAULT NULL COMMENT 'Period settlement status.   1 &#x3D; Event period is settled,  2 &#x3D; Event period is re-settled,  3 &#x3D; Event period is cancelled,  4 &#x3D; Event period is re-settled as cancelled,  5 &#x3D; Event is deleted ',
  `settlementId` BIGINT DEFAULT NULL COMMENT 'Unique id of the settlement. In case of a re-settlement, a new settlementId and settledAt will be generated.',
  `settledAt` DATETIME DEFAULT NULL COMMENT 'Date and time in UTC when the period was settled.',
  `team1Score` INT DEFAULT NULL COMMENT 'Team1 score.',
  `team2Score` INT DEFAULT NULL COMMENT 'Team2 score.',
  `team1ScoreSets` INT DEFAULT NULL COMMENT 'Team1 sets score. Supported for tennis only.',
  `team2ScoreSets` INT DEFAULT NULL COMMENT 'Team2 sets score. Supported for tennis only.',
  `cancellationReason` TEXT DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `SettledFixturesSport` generated from model 'SettledFixturesSport'
--

CREATE TABLE IF NOT EXISTS `SettledFixturesSport` (
  `sportId` INT DEFAULT NULL COMMENT 'Same as requested sport Id.',
  `last` BIGINT DEFAULT NULL COMMENT 'Use this value for the subsequent requests for since query parameter to get just the changes since previous response.',
  `leagues` TEXT DEFAULT NULL COMMENT 'Contains a list of Leagues.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `SettledSpecial` generated from model 'SettledSpecial'
-- Settled Special
--

CREATE TABLE IF NOT EXISTS `SettledSpecial` (
  `id` BIGINT DEFAULT NULL COMMENT 'Id for the Settled Special',
  `status` INT DEFAULT NULL COMMENT 'Status of the settled special.',
  `settlementId` BIGINT DEFAULT NULL COMMENT 'Id for the Settled Special',
  `settledAt` DATETIME DEFAULT NULL COMMENT 'Settled DateTime',
  `cancellationReason` TEXT DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Settled Special';

--
-- Table structure for table `SettledSpecialsLeague` generated from model 'SettledSpecialsLeague'
-- League Dto to hold all settled specials for the league
--

CREATE TABLE IF NOT EXISTS `SettledSpecialsLeague` (
  `id` INT DEFAULT NULL COMMENT 'League Id.',
  `specials` TEXT DEFAULT NULL COMMENT 'A collection of Settled Specials'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='League Dto to hold all settled specials for the league';

--
-- Table structure for table `SettledSpecialsResponse` generated from model 'SettledSpecialsResponse'
-- Response dto for SettledSpecials request
--

CREATE TABLE IF NOT EXISTS `SettledSpecialsResponse` (
  `sportId` INT DEFAULT NULL COMMENT 'Id of a sport for which to retrieve the odds.',
  `last` BIGINT DEFAULT NULL COMMENT 'Last index for the settled fixture',
  `leagues` TEXT DEFAULT NULL COMMENT 'List of Leagues.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Response dto for SettledSpecials request';

--
-- Table structure for table `SpecialBet` generated from model 'SpecialBet'
--

CREATE TABLE IF NOT EXISTS `SpecialBet` (
  `betId` BIGINT NOT NULL COMMENT 'Bet identification',
  `uniqueRequestId` TEXT DEFAULT NULL COMMENT 'Unique Request Id',
  `wagerNumber` INT NOT NULL COMMENT 'Wager identification. All bets placed thru the API will have value 1. Website Classic view supports multiple contest(special) bets placement in the same bet slip in that case the bet would have appropriate wager number, as well as all round robin parlay bets.',
  `placedAt` DATETIME NOT NULL COMMENT 'Date time when the bet was placed.',
  `betStatus` ENUM('ACCEPTED', 'CANCELLED', 'LOSE', 'REFUNDED', 'WON') NOT NULL COMMENT 'Bet Status.  ACCEPTED &#x3D; Bet was accepted,  CANCELLED &#x3D; Bet is cancelled as per Pinnacle betting rules,  LOSE &#x3D; The bet is settled as lose, REFUNDED &#x3D; When an event is cancelled or when the bet is settled as push, the bet will have REFUNDED status,  WON &#x3D; The bet is settled as won  ',
  `betType` TEXT NOT NULL,
  `win` DECIMAL(20, 9) NOT NULL COMMENT 'Win amount.',
  `risk` DECIMAL(20, 9) NOT NULL COMMENT 'Risk amount.',
  `winLoss` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Win-Loss for settled bets.',
  `oddsFormat` TEXT NOT NULL,
  `customerCommission` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Client’s commission on the bet.',
  `cancellationReason` TEXT DEFAULT NULL,
  `updateSequence` BIGINT NOT NULL COMMENT 'Update Sequence. It gets updated when the bet status change.',
  `specialId` BIGINT NOT NULL,
  `specialName` TEXT NOT NULL,
  `contestantId` BIGINT NOT NULL,
  `contestantName` TEXT NOT NULL,
  `price` DECIMAL(20, 9) NOT NULL,
  `handicap` DECIMAL(20, 9) DEFAULT NULL,
  `units` TEXT DEFAULT NULL,
  `sportId` INT NOT NULL,
  `leagueId` INT NOT NULL,
  `eventId` BIGINT DEFAULT NULL COMMENT 'Populated if bet was placed on a special linked to the event.',
  `periodNumber` INT DEFAULT NULL COMMENT 'Populated if bet was placed on a special linked to the event.',
  `team1` TEXT DEFAULT NULL COMMENT 'Populated if bet was placed on a special linked to the event.',
  `team2` TEXT DEFAULT NULL COMMENT 'Populated if bet was placed on a special linked to the event.',
  `eventStartTime` DATETIME DEFAULT NULL COMMENT 'Date time when the event starts'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `SpecialBetRequest` generated from model 'SpecialBetRequest'
--

CREATE TABLE IF NOT EXISTS `SpecialBetRequest` (
  `uniqueRequestId` TEXT DEFAULT NULL COMMENT 'This unique id of the place bet requests. This is to support idempotent requests.',
  `acceptBetterLine` TINYINT(1) DEFAULT NULL COMMENT 'Whether or not to accept a bet when there is a line change in favor of the client.',
  `oddsFormat` TEXT DEFAULT NULL,
  `stake` DECIMAL(20, 9) DEFAULT NULL COMMENT 'amount in client’s currency.',
  `winRiskStake` ENUM('WIN', 'RISK') DEFAULT NULL COMMENT 'Whether the stake amount is risk or win amount.',
  `lineId` BIGINT DEFAULT NULL COMMENT 'Line identification.',
  `specialId` BIGINT DEFAULT NULL COMMENT 'Special identification.',
  `contestantId` BIGINT DEFAULT NULL COMMENT 'Contestant identification.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `SpecialBetResponse` generated from model 'SpecialBetResponse'
--

CREATE TABLE IF NOT EXISTS `SpecialBetResponse` (
  `status` ENUM('ACCEPTED', 'PROCESSED_WITH_ERROR') DEFAULT NULL COMMENT 'Status of the request.',
  `errorCode` ENUM('ALL_BETTING_CLOSED', 'ABOVE_MAX_BET_AMOUNT', 'BELOW_MIN_BET_AMOUNT', 'BLOCKED_BETTING', 'BLOCKED_CLIENT', 'CONTEST_NOT_FOUND', 'DUPLICATE_UNIQUE_REQUEST_ID', 'INCOMPLETE_CUSTOMER_BETTING_PROFILE', 'INSUFFICIENT_FUNDS', 'INVALID_COUNTRY', 'INVALID_REQUEST', 'LINE_CHANGED', 'PAST_CUTOFFTIME', 'RESPONSIBLE_BETTING_LOSS_LIMIT_EXCEEDED', 'RESPONSIBLE_BETTING_RISK_LIMIT_EXCEEDED', 'SYSTEM_ERROR_1', 'SYSTEM_ERROR_2', 'UNIQUE_REQUEST_ID_REQUIRED', 'INVALID_CUSTOMER_PROFILE') DEFAULT NULL COMMENT 'When Status is PROCESSED_WITH_ERROR, provides a code indicating the specific problem.  ALL_BETTING_CLOSED &#x3D; Betting is not allowed at this moment. This may happen during system maintenance.    ABOVE_MAX_BET_AMOUNT &#x3D; Stake is above allowed maximum amount,    BELOW_MIN_BET_AMOUNT &#x3D; Stake is below allowed minimum amount,    BLOCKED_BETTING &#x3D; Betting is suspended for the client,    BLOCKED_CLIENT &#x3D; Client is no longer active,    CONTEST_NOT_FOUND &#x3D; Incorrect contest id provided or contest is no longer available,    DUPLICATE_UNIQUE_REQUEST_ID &#x3D; UniqueRequestId must be unique for each bet,    INCOMPLETE_CUSTOMER_BETTING_PROFILE &#x3D; Customer profile could not be loaded,     INSUFFICIENT_FUNDS &#x3D; Bet is submitted by a client with insufficient funds,    INVALID_COUNTRY &#x3D; Client country is not allowed for betting,    INVALID_REQUEST &#x3D; Special bet request is not valid,    LINE_CHANGED &#x3D; Bet is submitted on a line that has changed,    PAST_CUTOFFTIME &#x3D; Bet is submitted on a game after the betting cutoff time,    RESPONSIBLE_BETTING_LOSS_LIMIT_EXCEEDED &#x3D; Self-imposed loss limit exceeded,    RESPONSIBLE_BETTING_RISK_LIMIT_EXCEEDED &#x3D; Self-imposed risk limit exceeded,   SYSTEM_ERROR_1 &#x3D; Unexpected error,    SYSTEM_ERROR_2 &#x3D; Unexpected error,    UNIQUE_REQUEST_ID_REQUIRED &#x3D; UniqueRequestId is missing,    INVALID_CUSTOMER_PROFILE ',
  `betId` BIGINT DEFAULT NULL COMMENT 'Id of a newly created bet.',
  `uniqueRequestId` TEXT DEFAULT NULL COMMENT 'Unique identifier provided in the request.',
  `betterLineWasAccepted` TINYINT(1) DEFAULT NULL COMMENT 'Whether or not the bet was accepted on the line that changed in favour of client. This can be true only if acceptBetterLine in the Place Bet request is set to TRUE.',
  `specialBet` TEXT DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `SpecialFixture` generated from model 'SpecialFixture'
--

CREATE TABLE IF NOT EXISTS `SpecialFixture` (
  `id` BIGINT DEFAULT NULL COMMENT 'Unique Id',
  `betType` ENUM('MULTI_WAY_HEAD_TO_HEAD', 'SPREAD', 'OVER_UNDER') DEFAULT NULL COMMENT 'The type [MULTI_WAY_HEAD_TO_HEAD, SPREAD, OVER_UNDER]',
  `name` TEXT DEFAULT NULL COMMENT 'Name of the special.',
  `date` DATETIME DEFAULT NULL COMMENT 'Date of the special in UTC.',
  `cutoff` DATETIME DEFAULT NULL COMMENT 'Wagering cutoff date in UTC.',
  `category` TEXT DEFAULT NULL COMMENT 'The category that the special falls under.',
  `units` TEXT DEFAULT NULL COMMENT 'Measurment in the context of the special. This is applicable to specials bet type spead and over/under. In a hockey special this could be goals.',
  `status` ENUM('O', 'H', 'I') DEFAULT NULL COMMENT 'Status of the Special    O &#x3D; This is the starting status. It means that the lines  are open for betting,    H &#x3D; This status indicates that the lines are temporarily unavailable  for betting,    I &#x3D; This status indicates that one or more lines have a red circle  (a lower maximum bet amount) ',
  `event` TEXT DEFAULT NULL,
  `contestants` TEXT DEFAULT NULL COMMENT 'ContestantLines available for wagering.',
  `liveStatus` ENUM('0', '1', '2') DEFAULT NULL COMMENT 'When a special is linked to an event, we will return live status of the event, otherwise it will be 0.  0 &#x3D; No live betting will be offered on this event,  1 &#x3D; Live betting event,  2 &#x3D; Live betting will be offered on this match, but on a different event.   Please note that live delay is applied when placing bets on special with LiveStatus&#x3D;1  '
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `SpecialLineResponse` generated from model 'SpecialLineResponse'
--

CREATE TABLE IF NOT EXISTS `SpecialLineResponse` (
  `status` ENUM('SUCCESS', 'NOT_EXISTS') DEFAULT NULL COMMENT 'Status [SUCCESS &#x3D; OK, NOT_EXISTS &#x3D; Line not offered anymore]',
  `specialId` BIGINT DEFAULT NULL COMMENT 'Special Id.',
  `contestantId` BIGINT DEFAULT NULL COMMENT 'Contestant Id.',
  `minRiskStake` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Minimum bettable risk amount.',
  `maxRiskStake` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Maximum bettable risk amount.',
  `minWinStake` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Minimum bettable win amount.',
  `maxWinStake` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Maximum bettable win amount.',
  `lineId` BIGINT DEFAULT NULL COMMENT 'Line identification needed to place a bet.',
  `price` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Latest price.',
  `handicap` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Handicap.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `SpecialOddsContestantLine` generated from model 'SpecialOddsContestantLine'
--

CREATE TABLE IF NOT EXISTS `SpecialOddsContestantLine` (
  `id` BIGINT DEFAULT NULL COMMENT 'ContestantLine Id.',
  `lineId` BIGINT DEFAULT NULL COMMENT 'Line identifier required for placing a bet.',
  `price` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Price of the line.',
  `handicap` DECIMAL(20, 9) DEFAULT NULL COMMENT 'A number indicating the spread, over/under etc.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `SpecialOddsLeague` generated from model 'SpecialOddsLeague'
--

CREATE TABLE IF NOT EXISTS `SpecialOddsLeague` (
  `id` INT DEFAULT NULL COMMENT 'League id.',
  `specials` TEXT DEFAULT NULL COMMENT 'A collection of FixturesSpecial.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `SpecialOddsResponse` generated from model 'SpecialOddsResponse'
--

CREATE TABLE IF NOT EXISTS `SpecialOddsResponse` (
  `sportId` INT DEFAULT NULL COMMENT 'Id of a sport for which to retrieve the odds.',
  `last` BIGINT DEFAULT NULL COMMENT 'Used for retrieving changes only on subsequent requests. Provide this value as the Since paramter in subsequent calls to only retrieve changes.',
  `leagues` TEXT DEFAULT NULL COMMENT 'Contains a list of Leagues.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `SpecialOddsSpecial` generated from model 'SpecialOddsSpecial'
--

CREATE TABLE IF NOT EXISTS `SpecialOddsSpecial` (
  `id` BIGINT DEFAULT NULL COMMENT 'Special Id.',
  `maxBet` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Maximum bet volume amount. See [How to calculate max risk from the max volume](https://github.com/pinnacleapi/pinnacleapi-documentation/blob/master/FAQ.md#how-to-calculate-max-risk-from-the-max-volume-limits-in-odds)',
  `contestantLines` TEXT DEFAULT NULL COMMENT 'ContestantLines available for wagering on.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `SpecialsFixturesContestant` generated from model 'SpecialsFixturesContestant'
--

CREATE TABLE IF NOT EXISTS `SpecialsFixturesContestant` (
  `id` BIGINT DEFAULT NULL COMMENT 'Contestant Id.',
  `name` TEXT DEFAULT NULL COMMENT 'Name of the contestant.',
  `rotNum` INT DEFAULT NULL COMMENT 'Rotation Number.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `SpecialsFixturesEvent` generated from model 'SpecialsFixturesEvent'
-- Optional event asscoaited with the special.
--

CREATE TABLE IF NOT EXISTS `SpecialsFixturesEvent` (
  `id` INT DEFAULT NULL COMMENT 'Event Id',
  `periodNumber` INT DEFAULT NULL COMMENT 'The period of the match. For example in soccer 0 (Game), 1 (1st Half) &amp; 2 (2nd Half)',
  `home` TEXT DEFAULT NULL COMMENT 'Home team name.',
  `away` TEXT DEFAULT NULL COMMENT 'Away team name.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Optional event asscoaited with the special.';

--
-- Table structure for table `SpecialsFixturesLeague` generated from model 'SpecialsFixturesLeague'
--

CREATE TABLE IF NOT EXISTS `SpecialsFixturesLeague` (
  `id` INT DEFAULT NULL COMMENT 'FixturesLeague Id.',
  `specials` TEXT DEFAULT NULL COMMENT 'A collection of Specials'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `SpecialsFixturesResponse` generated from model 'SpecialsFixturesResponse'
--

CREATE TABLE IF NOT EXISTS `SpecialsFixturesResponse` (
  `sportId` INT DEFAULT NULL COMMENT 'Id of a sport for which to retrieve the odds.',
  `last` BIGINT DEFAULT NULL COMMENT 'Used for retrieving changes only on subsequent requests. Provide this value as the Since paramter in subsequent calls to only retrieve changes.',
  `leagues` TEXT DEFAULT NULL COMMENT 'Contains a list of Leagues.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `Sport` generated from model 'Sport'
--

CREATE TABLE IF NOT EXISTS `Sport` (
  `id` INT DEFAULT NULL COMMENT 'Sport Id.',
  `name` TEXT DEFAULT NULL COMMENT 'Sport name.',
  `hasOfferings` TINYINT(1) DEFAULT NULL COMMENT 'Whether the sport currently has events or specials.',
  `leagueSpecialsCount` INT DEFAULT NULL COMMENT 'Indicates how many specials are in the given sport.',
  `eventSpecialsCount` INT DEFAULT NULL COMMENT 'Indicates how many event specials are in the given sport.',
  `eventCount` INT DEFAULT NULL COMMENT 'Indicates how many events are in the given sport.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `SportPeriod` generated from model 'SportPeriod'
--

CREATE TABLE IF NOT EXISTS `SportPeriod` (
  `number` INT DEFAULT NULL COMMENT 'Period Number',
  `description` TEXT DEFAULT NULL COMMENT 'Description for the period',
  `shortDescription` TEXT DEFAULT NULL COMMENT 'Short description for the period',
  `spreadDescription` TEXT DEFAULT NULL COMMENT 'Description for the Spread',
  `moneylineDescription` TEXT DEFAULT NULL COMMENT 'Description for the Moneyline',
  `totalDescription` TEXT DEFAULT NULL COMMENT 'Description for the Totals',
  `team1TotalDescription` TEXT DEFAULT NULL COMMENT 'Description for Team1 Totals',
  `team2TotalDescription` TEXT DEFAULT NULL COMMENT 'Description for Team2 Totals',
  `spreadShortDescription` TEXT DEFAULT NULL COMMENT 'Short description for the Spread',
  `moneylineShortDescription` TEXT DEFAULT NULL COMMENT 'Short description for the Moneyline',
  `totalShortDescription` TEXT DEFAULT NULL COMMENT 'Short description for the Totals',
  `team1TotalShortDescription` TEXT DEFAULT NULL COMMENT 'Short description for Team1 Totals',
  `team2TotalShortDescription` TEXT DEFAULT NULL COMMENT 'Short description for Team2 Totals'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `SportsResponse` generated from model 'SportsResponse'
--

CREATE TABLE IF NOT EXISTS `SportsResponse` (
  `sports` TEXT DEFAULT NULL COMMENT 'Sports container.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `StraightBet` generated from model 'StraightBet'
--

CREATE TABLE IF NOT EXISTS `StraightBet` (
  `betId` BIGINT NOT NULL COMMENT 'Bet identification',
  `wagerNumber` INT NOT NULL COMMENT 'Wager identification. All bets placed thru the API will have value 1. Website Classic view supports multiple contest(special) bets placement in the same bet slip in that case the bet would have appropriate wager number, as well as all round robin parlay bets.',
  `placedAt` DATETIME NOT NULL COMMENT 'Date time when the bet was placed.',
  `betStatus` ENUM('ACCEPTED', 'CANCELLED', 'LOSE', 'PENDING_ACCEPTANCE', 'REFUNDED', 'NOT_ACCEPTED', 'WON') NOT NULL COMMENT 'Bet Status.    ACCEPTED &#x3D; Bet was accepted,   CANCELLED &#x3D; Bet is cancelled as per Pinnacle betting rules,   LOSE &#x3D; The bet is settled as lose,   PENDING_ACCEPTANCE &#x3D; This status is reserved only for live bets. If a live bet is placed during danger zone or live delay is applied, it will be in PENDING_ACCEPTANCE , otherwise in ACCEPTED status. From this status bet can go to ACCEPTED or NOT_ACCEPTED status,   REFUNDED &#x3D; When an event is cancelled or when the bet is settled as push, the bet will have REFUNDED status,   NOT_ACCEPTED &#x3D; Bet was not accepted. Bet can be in this status only if it was previously in PENDING_ACCEPTANCE status,   WON &#x3D; The bet is settled as won  ',
  `betType` ENUM('MONEYLINE', 'TEAM_TOTAL_POINTS', 'SPREAD', 'TOTAL_POINTS', 'SPECIAL', 'PARLAY', 'TEASER', 'MANUAL') NOT NULL COMMENT 'Bet type.',
  `win` DECIMAL(20, 9) NOT NULL COMMENT 'Win amount.',
  `risk` DECIMAL(20, 9) NOT NULL COMMENT 'Risk amount.',
  `winLoss` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Win-Loss for settled bets.',
  `oddsFormat` TEXT NOT NULL,
  `customerCommission` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Client’s commission on the bet.',
  `cancellationReason` TEXT DEFAULT NULL,
  `updateSequence` BIGINT NOT NULL COMMENT 'Update Sequence',
  `sportId` INT DEFAULT NULL,
  `leagueId` INT DEFAULT NULL,
  `eventId` BIGINT DEFAULT NULL,
  `handicap` DECIMAL(20, 9) DEFAULT NULL,
  `price` DECIMAL(20, 9) DEFAULT NULL,
  `teamName` TEXT DEFAULT NULL,
  `side` ENUM('OVER', 'UNDER') DEFAULT NULL COMMENT 'Side type.',
  `pitcher1` TEXT DEFAULT NULL COMMENT 'Pitcher name of team1. Only for bets on baseball.',
  `pitcher2` TEXT DEFAULT NULL COMMENT 'Pitcher name of team2. Only for bets on baseball.',
  `pitcher1MustStart` TEXT DEFAULT NULL COMMENT 'Whether the team1 pitcher must start. Only for bets on baseball.',
  `pitcher2MustStart` TEXT DEFAULT NULL COMMENT 'Whether the team1 pitcher must start. Only for bets on baseball.',
  `team1` TEXT DEFAULT NULL,
  `team2` TEXT DEFAULT NULL,
  `periodNumber` INT DEFAULT NULL,
  `team1Score` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Team 1 score that the bet was placed on, only for live bets.',
  `team2Score` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Team 2 score that the bet was placed, only for live bets.',
  `ftTeam1Score` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Full time team 1 score, only for settled bets.',
  `ftTeam2Score` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Full time team 2 score, only for settled bets.',
  `pTeam1Score` DECIMAL(20, 9) DEFAULT NULL COMMENT '.End of period team 1 score, only for settled bets. If the bet was placed on Game period (periodNumber &#x3D;0) , this will be null . ',
  `pTeam2Score` DECIMAL(20, 9) DEFAULT NULL COMMENT 'End of period team 2 score, only for settled bets. If the bet was placed on Game period (periodNumber &#x3D;0), this will be null',
  `isLive` TEXT DEFAULT NULL COMMENT 'Whether the bet is on live event'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `StraightBetV3` generated from model 'StraightBetV3'
--

CREATE TABLE IF NOT EXISTS `StraightBetV3` (
  `betId` BIGINT NOT NULL COMMENT 'Bet identification',
  `wagerNumber` INT NOT NULL COMMENT 'Wager identification. All bets placed thru the API will have value 1. Website Classic view supports multiple contest(special) bets placement in the same bet slip in that case the bet would have appropriate wager number, as well as all round robin parlay bets.',
  `placedAt` DATETIME NOT NULL COMMENT 'Date time when the bet was placed.',
  `betStatus` ENUM('ACCEPTED', 'CANCELLED', 'LOSE', 'PENDING_ACCEPTANCE', 'REFUNDED', 'NOT_ACCEPTED', 'WON') NOT NULL COMMENT 'Bet Status.    ACCEPTED &#x3D; Bet was accepted,   CANCELLED &#x3D; Bet is cancelled as per Pinnacle betting rules,   LOSE &#x3D; The bet is settled as lose,   PENDING_ACCEPTANCE &#x3D; This status is reserved only for live bets. If a live bet is placed during danger zone or live delay is applied, it will be in PENDING_ACCEPTANCE , otherwise in ACCEPTED status. From this status bet can go to ACCEPTED or NOT_ACCEPTED status,   REFUNDED &#x3D; When an event is cancelled or when the bet is settled as push, the bet will have REFUNDED status,   NOT_ACCEPTED &#x3D; Bet was not accepted. Bet can be in this status only if it was previously in PENDING_ACCEPTANCE status,   WON &#x3D; The bet is settled as won  ',
  `betType` ENUM('MONEYLINE', 'TEAM_TOTAL_POINTS', 'SPREAD', 'TOTAL_POINTS', 'SPECIAL', 'PARLAY', 'TEASER', 'MANUAL') NOT NULL COMMENT 'Bet type.',
  `win` DECIMAL(20, 9) NOT NULL COMMENT 'Win amount.',
  `risk` DECIMAL(20, 9) NOT NULL COMMENT 'Risk amount.',
  `winLoss` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Win-Loss for settled bets.',
  `oddsFormat` TEXT NOT NULL,
  `customerCommission` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Client’s commission on the bet.',
  `cancellationReason` TEXT DEFAULT NULL,
  `updateSequence` BIGINT NOT NULL COMMENT 'Update Sequence',
  `sportId` INT DEFAULT NULL,
  `leagueId` INT DEFAULT NULL,
  `eventId` BIGINT DEFAULT NULL,
  `handicap` DECIMAL(20, 9) DEFAULT NULL,
  `price` DECIMAL(20, 9) DEFAULT NULL,
  `teamName` TEXT DEFAULT NULL,
  `side` ENUM('OVER', 'UNDER') DEFAULT NULL COMMENT 'Side type.',
  `pitcher1` TEXT DEFAULT NULL COMMENT 'Pitcher name of team1. Only for bets on baseball.',
  `pitcher2` TEXT DEFAULT NULL COMMENT 'Pitcher name of team2. Only for bets on baseball.',
  `pitcher1MustStart` TINYINT(1) DEFAULT NULL COMMENT 'Baseball only. Refers to the pitcher for Team1.  This applicable only for MONEYLINE bet type, for all other bet types this has to be TRUE.',
  `pitcher2MustStart` TINYINT(1) DEFAULT NULL COMMENT 'Baseball only. Refers to the pitcher for Team2.  This applicable only for MONEYLINE bet type, for all other bet types this has to be TRUE.',
  `team1` TEXT DEFAULT NULL,
  `team2` TEXT DEFAULT NULL,
  `periodNumber` INT DEFAULT NULL,
  `team1Score` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Team 1 score that the bet was placed on, only for live bets.',
  `team2Score` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Team 2 score that the bet was placed, only for live bets.',
  `ftTeam1Score` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Full time team 1 score, only for settled bets.',
  `ftTeam2Score` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Full time team 2 score, only for settled bets.',
  `pTeam1Score` DECIMAL(20, 9) DEFAULT NULL COMMENT '.End of period team 1 score, only for settled bets. If the bet was placed on Game period (periodNumber &#x3D;0), this will be null . ',
  `pTeam2Score` DECIMAL(20, 9) DEFAULT NULL COMMENT 'End of period team 2 score, only for settled bets. If the bet was placed on Game period (periodNumber &#x3D;0), this will be null',
  `isLive` TINYINT(1) DEFAULT NULL COMMENT 'Whether the bet is on live event',
  `eventStartTime` DATETIME DEFAULT NULL COMMENT 'Date time when the event starts.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `SuccessfulCurrenciesResponse` generated from model 'SuccessfulCurrenciesResponse'
--

CREATE TABLE IF NOT EXISTS `SuccessfulCurrenciesResponse` (
  `currencies` TEXT DEFAULT NULL COMMENT 'Currencies container.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `TeaserBet` generated from model 'TeaserBet'
--

CREATE TABLE IF NOT EXISTS `TeaserBet` (
  `betId` BIGINT NOT NULL COMMENT 'Bet identification',
  `uniqueRequestId` TEXT DEFAULT NULL COMMENT 'Unique Request Id',
  `wagerNumber` INT NOT NULL COMMENT 'Wager identification. All bets placed thru the API will have value 1. Website Classic view supports multiple contest(special) bets placement in the same bet slip in that case the bet would have appropriate wager number, as well as all round robin parlay bets.',
  `placedAt` DATETIME NOT NULL COMMENT 'Date time when the bet was placed.',
  `betStatus` ENUM('ACCEPTED', 'CANCELLED', 'LOSE', 'REFUNDED', 'WON') NOT NULL COMMENT 'Bet Status.   ACCEPTED &#x3D; Bet was accepted,   CANCELLED &#x3D; Bet is cancelled as per Pinnacle betting rules,   LOSE &#x3D; The bet is settled as lose,   REFUNDED &#x3D; When an event is cancelled or when the bet is settled as push, the bet will have REFUNDED status,   WON &#x3D; The bet is settled as won  ',
  `betType` TEXT NOT NULL,
  `win` DECIMAL(20, 9) NOT NULL COMMENT 'Win amount.',
  `risk` DECIMAL(20, 9) NOT NULL COMMENT 'Risk amount.',
  `winLoss` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Win-Loss for settled bets.',
  `oddsFormat` TEXT NOT NULL,
  `customerCommission` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Client’s commission on the bet.',
  `cancellationReason` TEXT DEFAULT NULL,
  `updateSequence` BIGINT NOT NULL COMMENT 'Update Sequence',
  `teaserName` TEXT NOT NULL,
  `isSameEventOnly` TINYINT(1) NOT NULL,
  `minPicks` DECIMAL(20, 9) NOT NULL,
  `maxPicks` DECIMAL(20, 9) NOT NULL,
  `price` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Populated for all teaser bets and will be the original price at the time of the placement.',
  `finalPrice` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Only for settled parlay. Final price may differ in case leg was cancelled or half won.',
  `teaserId` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Reference to the teaser id',
  `teaserGroupId` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Reference to the teaser group id',
  `legs` TEXT NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `TeaserGroups` generated from model 'TeaserGroups'
--

CREATE TABLE IF NOT EXISTS `TeaserGroups` (
  `id` BIGINT DEFAULT NULL COMMENT 'Unique identifier.',
  `name` TEXT DEFAULT NULL COMMENT 'Friendly name for the Teaser Group',
  `teasers` TEXT DEFAULT NULL COMMENT 'A collection of Teaser.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `TeaserGroupsBetType` generated from model 'TeaserGroupsBetType'
--

CREATE TABLE IF NOT EXISTS `TeaserGroupsBetType` (
  `points` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Number of points the line will be teased for the given league.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `TeaserGroupsLeague` generated from model 'TeaserGroupsLeague'
--

CREATE TABLE IF NOT EXISTS `TeaserGroupsLeague` (
  `id` INT DEFAULT NULL COMMENT 'Unique identifier. League details can be retrieved from a call to v2/leagues endpoint.',
  `spread` TEXT DEFAULT NULL,
  `total` TEXT DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `TeaserGroupsPayout` generated from model 'TeaserGroupsPayout'
--

CREATE TABLE IF NOT EXISTS `TeaserGroupsPayout` (
  `numberOfLegs` INT DEFAULT NULL COMMENT 'Number of legs that must be bet and won to get the associated price.',
  `price` DECIMAL(20, 9) DEFAULT NULL COMMENT 'Price of the bet given the specified number of legs.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `TeaserGroupsResponse` generated from model 'TeaserGroupsResponse'
--

CREATE TABLE IF NOT EXISTS `TeaserGroupsResponse` (
  `teaserGroups` TEXT DEFAULT NULL COMMENT 'A collection of TeaserGroups containing available teasers.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `TeaserGroupsTeaser` generated from model 'TeaserGroupsTeaser'
--

CREATE TABLE IF NOT EXISTS `TeaserGroupsTeaser` (
  `id` BIGINT DEFAULT NULL COMMENT 'Unique identifier.',
  `description` TEXT DEFAULT NULL COMMENT 'Description for the Teaser.',
  `sportId` INT DEFAULT NULL COMMENT 'Unique Sport identifier. Sport details can be retrieved from a call to v2/sports endpoint.',
  `minLegs` INT DEFAULT NULL COMMENT 'Minimum number of legs that must be selected.',
  `maxLegs` INT DEFAULT NULL COMMENT 'Maximum number of legs that can be selected.',
  `sameEventOnly` TINYINT(1) DEFAULT NULL COMMENT 'If &#39;true&#39; then all legs must be from the same event, otherwise legs can be from different events.',
  `payouts` TEXT DEFAULT NULL COMMENT 'A collection of Payout indicating all possible payout combinations.',
  `leagues` TEXT DEFAULT NULL COMMENT 'A collection of Leagues available to the teaser.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `TeaserLeg` generated from model 'TeaserLeg'
--

CREATE TABLE IF NOT EXISTS `TeaserLeg` (
  `sportId` INT DEFAULT NULL,
  `legBetType` ENUM('SPREAD', 'TOTAL_POINTS') DEFAULT NULL COMMENT 'Teaser leg type.',
  `legBetStatus` ENUM('CANCELLED', 'LOSE', 'PUSH', 'REFUNDED', 'WON') DEFAULT NULL COMMENT 'CANCELLED &#x3D; The leg is canceled- the stake on this leg will be transferred to the next one. In this case the leg will be ignored when calculating the winLoss,   LOSE &#x3D; The leg is a loss or a push-lose. When Push-lose happens, the half of the stake on the leg will be pushed to the next leg, and the other half will be a lose. This can happen only when the leg is placed on a quarter points handicap,   PUSH &#x3D; The leg is a push - the stake on this leg will be transferred to the next one. In this case the leg will be ignored when calculating the winLoss,   REFUNDED &#x3D; The leg is refunded - the stake on this leg will be transferred to the next one. In this case the leg will be ignored when calculating the winLoss,   WON &#x3D; The leg is a won or a push-won. When Push-won happens, the half of the stake on the leg will be pushed to the next leg, and the other half is won. This can happen only when the leg is placed on a quarter points handicap   ',
  `leagueId` INT DEFAULT NULL,
  `eventId` BIGINT DEFAULT NULL,
  `eventStartTime` DATETIME DEFAULT NULL COMMENT 'Date time when the event starts.',
  `handicap` DECIMAL(20, 9) DEFAULT NULL,
  `teamName` TEXT DEFAULT NULL,
  `side` ENUM('OVER', 'UNDER') DEFAULT NULL COMMENT 'Side type.',
  `team1` TEXT DEFAULT NULL,
  `team2` TEXT DEFAULT NULL,
  `periodNumber` INT DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `Translation` generated from model 'Translation'
--

CREATE TABLE IF NOT EXISTS `Translation` (
  `text` TEXT DEFAULT NULL COMMENT 'Original requested text to be translated.',
  `cultures` TEXT DEFAULT NULL COMMENT 'Collection of translations by culture.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `TranslationByCulture` generated from model 'TranslationByCulture'
--

CREATE TABLE IF NOT EXISTS `TranslationByCulture` (
  `id` ENUM('en-US', 'en-GB', 'zh-CN', 'zh-TW', 'fi-FI', 'de-DE', 'he-IL', 'it-IT', 'nb-NO', 'pt-BR', 'ru-RU', 'es-ES', 'sv-SE', 'th-TH', 'pl-PL', 'fr-FR', 'ja-JP', 'ko-KR', 'vi-VN', 'id-ID', 'cs-CZ') DEFAULT NULL COMMENT 'Culture based on which the text is translated.',
  `text` TEXT DEFAULT NULL COMMENT 'Translation text.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Table structure for table `TranslationResponse` generated from model 'TranslationResponse'
--

CREATE TABLE IF NOT EXISTS `TranslationResponse` (
  `translations` TEXT DEFAULT NULL COMMENT 'Collection of translations.'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

